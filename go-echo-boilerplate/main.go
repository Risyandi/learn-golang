package main

import (
	"boilerplate/config"
	initDB "boilerplate/init"
	"boilerplate/module/auth"
	"boilerplate/schema/response"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Config
	config.LoadConfig()

	// Init DB
	initDB.InitMongoDB(config.AppConfig.MongoDBURI, config.AppConfig.DatabaseName)

	e := echo.New()

	// Middleware
	// Global middleware
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.BodyLimit("1M"))
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())

	// Register routes
	authGroup := e.Group("/api/v1")
	auth.RegisterRoutes(authGroup, config.AppConfig)

	router := e.Group("/health")

	// Health check
	router.GET("", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status":  "OK",
			"message": "Server is running",
		})
	})

	router.Any("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response.Base{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
			Message:    http.StatusText(http.StatusOK),
			Timestamp:  time.Now(),
			Data:       nil,
		})
	})

	e.Logger.Fatal(e.Start(":" + config.AppConfig.Port))
}
