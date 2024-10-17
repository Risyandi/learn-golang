package middleware

import (
	"bytes"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/config"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/constants"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/database/query/session"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/exception"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/internal/repository"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/logger"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/pkg/utils"
)

type (
	SessionMiddleware interface {
		CreateSession() fiber.Handler
		LogSession() fiber.Handler
	}

	sessionMiddleware struct {
		conf        *config.Config
		sessionRepo repository.SessionRepository
		log         zerolog.Logger
		exception   exception.Exception
	}
)

func NewSessionMiddleware(sessionRepo repository.SessionRepository) SessionMiddleware {
	return &sessionMiddleware{
		conf:        config.Get(),
		sessionRepo: sessionRepo,
		log:         logger.Get("session-middleware"),
		exception:   exception.NewException("session-middleware"),
	}
}

func (m sessionMiddleware) CreateSession() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userInfo := utils.CtxUserInfo(ctx.Context())
		session := &session.UserSessionEntity{
			SessionID: userInfo[constants.ClaimsTokenID].(string),
			UserID:    userInfo[constants.ClaimsSub].(string),
			IPAddress: ctx.IP(),
			URI:       ctx.OriginalURL(),
			UserAgent: string(ctx.Context().UserAgent()),
			IsActive:  true,
			CreatedAt: utils.GetLocalDateTime(),
			ExpiredAt: utils.GetLocalDateTime().AddDate(0, 1, 0),
		}
		day := utils.ResetDayTime(utils.GetLocalDateTime())
		_, err := m.sessionRepo.GetActiveSession(ctx.Context(), session.UserID, day)
		if err == nil {
			return ctx.Next()
		}
		tx, err := m.sessionRepo.NewTransaction(ctx.Context())

		err = tx.InsertSession(session)
		if err != nil {
			m.log.Error().Err(err).Msg("Insert session error")
		}

		// Temporary  comment, update session state is not needed yet
		// _, err = tx.DisableExpiredSession(session.UserID, day)

		tx.Commit()

		return ctx.Next()
	}
}

func (m sessionMiddleware) LogSession() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userInfo := utils.CtxUserInfo(ctx.Context())

		tx, err := m.sessionRepo.NewTransaction(ctx.Context())

		// Logging response body
		buffer := &bytes.Buffer{}
		resBody := buffer.Bytes()
		logs := session.UserLogEntity{
			SessionID: userInfo[constants.ClaimsTokenID].(string),
			UserID:    userInfo[constants.ClaimsSub].(string),
			Method:    ctx.Method(),
			URI:       ctx.OriginalURL(),
			IPAddress: ctx.IP(),
			UserAgent: string(ctx.Context().UserAgent()),
			Request:   utils.ByteToString(ctx.Body()),
			Response:  utils.ByteToString(resBody),
			CreatedAt: utils.GetLocalDateTime(),
		}

		err = tx.InsertLog("user_logs", logs)
		if err != nil {
			m.log.Error().Any("log", logs).Err(err).Msg("Insert user log error")
		}

		// admin log
		if userInfo[constants.ClaimsOAID] != nil {
			admin := session.AdminLogEntity{
				UserId:    userInfo[constants.ClaimsSub].(string),
				AdminId:   userInfo[constants.ClaimsOAID].(string),
				IP:        ctx.IP(),
				Method:    ctx.Method(),
				Uri:       ctx.OriginalURL(),
				UserAgent: string(ctx.Context().UserAgent()),
				Request:   utils.ByteToString(ctx.Body()),
				CreatedAt: utils.GetLocalDateTime(),
			}

			err = tx.InsertLog("admin_logs", admin)
			if err != nil {
				m.log.Error().Err(err).Msg("Insert admin log error")
			}
		}

		tx.Commit()

		return ctx.Next()
	}
}
