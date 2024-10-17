package exception

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"runtime"

	"github.com/rs/zerolog"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/constants"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/internal/response"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/logger"
)

type Exception interface {
	Error(err error)
	ErrorWithMessage(err error, message string)
	ErrorWithoutNoSqlResult(err error)
	BadRequest(err error)
	BadRequestMessage(err error, message string)
	NotFound(isList bool, modules ...string)
	Forbidden(message string)
	Unauthorized(message string)
	Conflict(message string)
	Unprocessable(message string)
}

type exception struct {
	log *zerolog.Logger
}

func NewException(types string) Exception {
	return &exception{
		log: logger.GetWithoutCaller(types),
	}
}

func (e *exception) getCaller() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d", file, line)

}

func (e *exception) defaultData(isList bool) interface{} {
	if isList {
		return []any{}
	}
	return nil
}

func (e *exception) Error(err error) {
	if err != nil {
		// e.log.Error().Str("caller", e.getCaller()).Err(err).Msg("SERVER ERROR")
		panic(response.NewErrorMessage(http.StatusInternalServerError, constants.ResponseErrorMessage, nil))
	}
}

func (e *exception) ErrorWithoutNoSqlResult(err error) {
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		// e.log.Error().Str("caller", e.getCaller()).Err(err).Msg("SERVER ERROR")
		panic(response.NewErrorMessage(http.StatusInternalServerError, constants.ResponseErrorMessage, nil))
	}
}

func (e *exception) BadRequest(err error) {
	if err != nil {
		// e.log.Error().Str("caller", e.getCaller()).Err(err).Msg("CLIENT ERROR")
		panic(response.NewErrorMessage(http.StatusBadRequest, err.Error(), nil))
	}
}

func (e *exception) BadRequestMessage(err error, message string) {
	if err != nil {
		// e.log.Error().Str("caller", e.getCaller()).Err(err).Msg("CLIENT ERROR")
		panic(response.NewErrorMessage(http.StatusBadRequest, message, nil))
	}
}

func (e *exception) Unprocessable(message string) {
	// e.log.Error().Str("caller", e.getCaller()).Msg("UNPROCESSABLE ENTITY")
	panic(response.NewErrorMessage(http.StatusUnprocessableEntity, message, nil))
}

func (e *exception) Unauthorized(message string) {
	// e.log.Error().Str("caller", e.getCaller()).Msg("NOT AUTHORIZED")
	panic(response.NewErrorMessage(http.StatusUnauthorized, message, nil))
}

func (e *exception) ErrorWithMessage(err error, message string) {
	if err != nil {
		// e.log.Error().Str("caller", e.getCaller()).Err(err).Msg("SERVER ERROR")
		panic(response.NewErrorMessage(http.StatusInternalServerError, message, nil))
	}
}

func (e *exception) NotFound(isList bool, modules ...string) {

	data := e.defaultData(isList)
	// e.log.Info().Str("caller", e.getCaller()).Msg("NOT FOUND")
	panic(response.NewErrorMessage(http.StatusNotFound, constants.ResponseErrorNotFound, data))
}

func (e *exception) Forbidden(message string) {
	// e.log.Error().Str("caller", e.getCaller()).Msg("FORBIDDEN")
	panic(response.NewErrorMessage(http.StatusForbidden, message, nil))
}

func (e *exception) Conflict(message string) {
	// e.log.Error().Str("caller", e.getCaller()).Msg("CONFLICT")
	panic(response.NewErrorMessage(http.StatusConflict, message, nil))
}
