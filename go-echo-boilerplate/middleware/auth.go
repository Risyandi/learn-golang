package middleware

import (
	"boilerplate/constant"
	"boilerplate/pkg/helper"
	"boilerplate/pkg/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func JWTAuth(jwtHelper *helper.JWTHelper) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get(constant.AuthHeaderKey)

			if authHeader == "" {
				return utils.ErrorResponse(c, http.StatusUnauthorized, "missing authorization header")
			}

			if !strings.HasPrefix(authHeader, constant.AuthBearerPrefix) {
				return utils.ErrorResponse(c, http.StatusUnauthorized, "invalid authorization header format")
			}

			tokenString := strings.TrimPrefix(authHeader, constant.AuthBearerPrefix)

			claims, err := jwtHelper.ValidateToken(tokenString)
			if err != nil {
				return utils.ErrorResponse(c, http.StatusUnauthorized, "invalid token")
			}

			c.Set(constant.AuthUserKey, claims)
			return next(c)
		}
	}
}

func AdminOnly() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims, ok := c.Get(constant.AuthUserKey).(*helper.JWTClaims)
			if !ok {
				return utils.ErrorResponse(c, http.StatusUnauthorized, "unauthorized")
			}

			if claims.Role != constant.RoleAdmin {
				return utils.ErrorResponse(c, http.StatusForbidden, "admin access required")
			}

			return next(c)
		}
	}
}
