package main

import (
	"crud-jne/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/users", handler.CreateUsers)
	e.GET("/users", handler.GetUsers)
	e.GET("/users/:id", handler.GetUser)
	e.PUT("/users/:id", handler.UpdateUser)
	e.DELETE("/users/:id", handler.DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}
