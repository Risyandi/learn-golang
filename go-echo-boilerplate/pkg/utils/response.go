package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type Meta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

func SuccessResponse(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, BaseResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func SuccessResponseWithMeta(c echo.Context, message string, data interface{}, meta *Meta) error {
	return c.JSON(http.StatusOK, BaseResponse{
		Status:  "success",
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func ErrorResponse(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, BaseResponse{
		Status:  "error",
		Message: message,
	})
}
