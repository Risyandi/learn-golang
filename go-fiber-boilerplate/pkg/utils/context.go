package utils

import (
	"context"

	"github.com/golang-jwt/jwt"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/constants"
)

func CtxUserInfo(ctx context.Context) jwt.MapClaims {
	if claims, ok := ctx.Value(constants.JwtClaimsKey).(jwt.MapClaims); ok {
		return claims
	}
	return nil
}
