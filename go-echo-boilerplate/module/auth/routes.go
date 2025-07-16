package auth

import (
	"github.com/labstack/echo/v4"

	"boilerplate/config"
)

// RegisterRoutes sets up the authentication-related routes.
func RegisterRoutes(e *echo.Group, config *config.Config) {
	authHandler := NewAuthHandler(config)
	authGroup := e.Group("/auth")
	authGroup.POST("/register", authHandler.RegisterHandler)
	authGroup.GET("/users", authHandler.GetUsersHandler)
	authGroup.GET("/users/:id", authHandler.GetUserByIDHandler)
	authGroup.POST("/login", authHandler.LoginHandler)
	authGroup.PUT("/users/:id", authHandler.UpdateUserHandler)
}
