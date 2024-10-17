package exception

import (
	"errors"
	"net/http"

	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/internal/response"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/pkg/utils"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	status := http.StatusText(code)
	var errorMessage string

	var fiberError *fiber.Error
	var data interface{} = err.Error()

	if errors.As(err, &fiberError) {

		code = fiberError.Code
		status = http.StatusText(code)
		errorMessage = err.Error()
		data = nil
	}

	var httpError *response.BaseResponse
	if errors.As(err, &httpError) {
		code = httpError.StatusCode
		data = httpError.Data
		status = http.StatusText(code)
		if httpError.Status != "" {
			status = httpError.Status
		}

		if httpError.Message != "" {
			translated, err := fiberi18n.Localize(ctx, httpError.Message)
			if err != nil {
				errorMessage = httpError.Message
			} else {

				errorMessage = translated
			}
		} else {
			errorMessage = http.StatusText(code)
		}
	}

	return ctx.Status(code).JSON(response.BaseResponse{
		Status:     status,
		StatusCode: code,
		Message:    errorMessage,
		Timestamp:  utils.GetLocalDateTime(),
		Data:       data,
	})
}
