package handler

import (
	"crud-jne/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var users = []model.User{}
var idCount = 1

func CreateUsers(c echo.Context) error {
	user := new(model.User)

	if err := c.Bind(user); err != nil {
		return err
	}

	user.ID = idCount
	idCount++
	users = append(users, *user)
	return c.JSON(http.StatusCreated, users)
}

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, u := range users {
		if u.ID == id {
			return c.JSON(http.StatusOK, u)
		}
	}
	return c.NoContent(http.StatusNotFound)
}

func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	updated := new(model.User)
	if err := c.Bind(updated); err != nil {
		return err
	}
	for i, u := range users {
		if u.ID == id {
			users[i].Name = updated.Name
			return c.JSON(http.StatusOK, users[i])
		}
	}
	return c.NoContent(http.StatusNotFound)
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.NoContent(http.StatusNotFound)
}
