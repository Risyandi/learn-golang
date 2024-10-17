package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/config"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/exception"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/logger"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/pkg/utils"
)

type (
	RequestMiddleware interface {
		Sanitize() fiber.Handler
	}

	requestMiddleware struct {
		conf      *config.Config
		log       zerolog.Logger
		exception exception.Exception
	}
)

func NewRequestMiddleware() RequestMiddleware {
	return &requestMiddleware{
		conf:      config.Get(),
		log:       logger.Get("request-middleware"),
		exception: exception.NewException("request-middleware"),
	}
}

func (m requestMiddleware) Sanitize() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Get the query parameters from the request
		queryParams := ctx.Queries()
		// Iterate through each query parameter and sanitize them
		for key, value := range queryParams {
			if value != "" {
				// Sanitize the parameter value
				sanitizedValue, err := utils.Sanitize(value)
				if err != nil {
					m.exception.BadRequest(err)
				}
				// Store the sanitized value in the list
				ctx.Request().URI().QueryArgs().Add(key, sanitizedValue)
			}
		}

		// Call the next handler
		return ctx.Next()
	}
}
