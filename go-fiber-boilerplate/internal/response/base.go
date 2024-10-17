package response

import (
	"net/http"
	"time"

	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/pkg/utils"
)

type BaseResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"statusCode" example:"200"`
	Message    string      `json:"message" example:"this is success"`
	Timestamp  time.Time   `json:"timestamp"`
	Data       interface{} `json:"data"`
}

func (res *BaseResponse) Error() string {
	return res.Message
}

func NewErrorMessage(code int, message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		StatusCode: code,
		Message:    message,
		Data:       data,
	}
}

func OK(ctx *fiber.Ctx, message string, data interface{}) error {
	return ctx.JSON(&BaseResponse{
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
		Message:    fiberi18n.MustLocalize(ctx, message),
		Timestamp:  utils.GetLocalDateTime(),
		Data:       data,
	})
}

func Created(ctx *fiber.Ctx, message string, data interface{}) error {
	return ctx.JSON(&BaseResponse{
		Status:     http.StatusText(http.StatusCreated),
		StatusCode: http.StatusCreated,
		Message:    fiberi18n.MustLocalize(ctx, message),
		Timestamp:  utils.GetLocalDateTime(),
		Data:       data,
	})
}

func Accepted(ctx *fiber.Ctx, message string, data interface{}) error {
	return ctx.JSON(&BaseResponse{
		Status:     http.StatusText(http.StatusAccepted),
		StatusCode: http.StatusAccepted,
		Message:    fiberi18n.MustLocalize(ctx, message),
		Timestamp:  utils.GetLocalDateTime(),
		Data:       data,
	})
}

func NewCustomMessage(code int, message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		Status:     http.StatusText(code),
		StatusCode: code,
		Message:    message,
		Timestamp:  utils.GetLocalDateTime(),
		Data:       data,
	}
}
