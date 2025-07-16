package auth

import (
	"boilerplate/config"
	"boilerplate/pkg/utils"
	"boilerplate/schema/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService *AuthService
}

func NewAuthHandler(config *config.Config) *AuthHandler {
	return &AuthHandler{
		authService: NewAuthService(config),
	}
}

func (h *AuthHandler) RegisterHandler(c echo.Context) error {
	req := new(request.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	user, err := h.authService.RegisterUser(c.Request().Context(), req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, err.Error())
	}

	return utils.SuccessResponse(c, "user registered successfully", user)
}

func (h *AuthHandler) GetUsersHandler(c echo.Context) error {
	// Extract pagination params from query
	page, limit := utils.GetPaginationParams(c)

	users, err := h.authService.GetUsers(c.Request().Context(), page, limit)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, err.Error())
	}

	return utils.SuccessResponse(c, "users fetched successfully", users)
}

func (h *AuthHandler) LoginHandler(c echo.Context) error {
	req := new(request.LoginRequest)
	if err := c.Bind(req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	user, err := h.authService.AuthenticateUser(c.Request().Context(), req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, err.Error())
	}

	return utils.SuccessResponse(c, "user logged in successfully", user)
}

func (h *AuthHandler) GetUserByIDHandler(c echo.Context) error {
	userID := c.Param("id")
	user, err := h.authService.GetUserByID(c.Request().Context(), userID)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, err.Error())
	}

	return utils.SuccessResponse(c, "user fetched successfully", user)
}

func (h *AuthHandler) UpdateUserHandler(c echo.Context) error {
	req := new(request.UpdateUserRequest)
	if err := c.Bind(req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err := h.authService.UpdateUser(ctx, req.Email, req)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, err.Error())
	}

	return utils.SuccessResponse(c, "user updated successfully", nil)
}
