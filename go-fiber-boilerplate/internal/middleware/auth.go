package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/config"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/constants"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/database"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/exception"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/internal/repository"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/logger"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/pkg/utils"
)

type (
	AuthMiddleware interface {
		Auth() fiber.Handler
		ValidateAccount() fiber.Handler
		ValidateToken() fiber.Handler
	}

	authMiddleware struct {
		conf      *config.Config
		authRepo  repository.AuthRepository
		redis     database.RedisInstance
		log       zerolog.Logger
		exception exception.Exception
	}

	AuthToken struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refreshToken"`
	}

	AuthStatus struct {
		IsActive    bool `json:"isActive"`
		IsSuspended bool `json:"isSuspended"`
	}
)

func NewAuthMiddleware(authRepo repository.AuthRepository, redis database.RedisInstance) AuthMiddleware {
	return &authMiddleware{
		conf:      config.Get(),
		authRepo:  authRepo,
		redis:     redis,
		log:       logger.Get("auth-middleware"),
		exception: exception.NewException("auth-middleware"),
	}
}

func (m *authMiddleware) Auth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")

		if authHeader == "" {
			m.exception.Unauthorized("Token is required")
		}
		token := authHeader[7:] //"Bearer "

		claims := jwt.MapClaims{}
		jwtToken, err := jwt.ParseWithClaims(
			token,
			claims,
			func(t *jwt.Token) (i interface{}, err error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return
				}

				key := m.conf.Encryption.AccessTokenPassphrase
				i = []byte(key)

				return
			},
		)

		if err != nil {
			m.exception.Unauthorized(err.Error())
		}

		//Check if is expired
		clm, ok := jwtToken.Claims.(jwt.MapClaims)
		if !ok || !jwtToken.Valid {
			m.exception.Unauthorized("Invalid or expired")
		} else {
			ctx.Locals(constants.JwtClaimsKey, clm)

		}

		return ctx.Next()
	}
}

func (m *authMiddleware) ValidateAccount() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var (
			status   = &AuthStatus{}
			userInfo = utils.CtxUserInfo(ctx.Context())
			userId   = userInfo[constants.ClaimsSub].(string)
		)

		_ = m.redis.Get(userId+"_status", status)

		if status == (&AuthStatus{}) {
			userStatus, _ := m.authRepo.GetUserStatusRepository(ctx.Context(), uuid.MustParse(userId))
			status.IsActive = userStatus.IsActive
			status.IsSuspended = userStatus.IsSuspended
			_ = m.redis.Setx(userId+"_status", status)

		}
		if status.IsSuspended {
			m.exception.Unauthorized(constants.ErrAccountSuspended)
		}
		return ctx.Next()
	}
}

func (m *authMiddleware) ValidateToken() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userInfo := utils.CtxUserInfo(ctx.Context())
		userId := userInfo[constants.ClaimsSub].(string)
		tokenId := userInfo[constants.ClaimsTokenID].(string)

		t := new(AuthToken)
		keys := fmt.Sprintf("%s_%s", userId, tokenId)
		err := m.redis.Get(keys, t)

		if err != nil || t.Token == "" {
			m.exception.Unauthorized(constants.ErrInvalidOrEmptyToken)
		}

		return ctx.Next()
	}
}
