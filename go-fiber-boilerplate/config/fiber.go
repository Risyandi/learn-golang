package config

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
)

func FiberZerolog(rawLogs *zerolog.Logger, logs zerolog.Logger) fiberzerolog.Config {
	// field logging will can see in the terminal
	fields := []string{
		fiberzerolog.FieldRequestID,
		fiberzerolog.FieldIP,
		fiberzerolog.FieldMethod,
		fiberzerolog.FieldStatus,
		fiberzerolog.FieldURL,
		fiberzerolog.FieldPath,
		fiberzerolog.FieldQueryParams,
		fiberzerolog.FieldBody,
		// fiberzerolog.FieldResBody,
		fiberzerolog.FieldLatency,
		fiberzerolog.FieldUserAgent,
		fiberzerolog.FieldError,
		fiberzerolog.FieldBytesReceived,
		fiberzerolog.FieldBytesSent,
	}

	// set configuration fiberzerolog
	return fiberzerolog.Config{
		Logger: &logs,
		Fields: fields,
		SkipBody: func(ctx *fiber.Ctx) bool {
			return strings.Contains(string(ctx.Request().Header.ContentType()), "multipart/form-data")
		},
		GetLogger: func(c *fiber.Ctx) zerolog.Logger {
			call, ok := c.Locals("caller").(string)
			if !ok {
				return logs
			}

			return rawLogs.With().Str("caller", call).Logger()
		},
		WrapHeaders: true,
		// GetResBody: func(c *fiber.Ctx) []byte {
		// 	return c.Body()
		// },
	}
}

func FiberRecover() recover.Config {
	// setup recovers from panics anywhere in the stack chain
	return recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			_, file, line, ok := runtime.Caller(4)
			if !ok {
				return
			}
			c.Locals("caller", fmt.Sprintf("%s:%d", file, line))
		},
	}
}
