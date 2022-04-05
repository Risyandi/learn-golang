package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// initialize the router
	router := gin.Default()

	// group version 1
	groupApiVersionOne := router.Group("/v1")
	groupApiVersionOne.GET("users", func(c *gin.Context) {
		c.JSON(http.StatusOK, "List Of group version one Users")
	})

	// User only can be added by authorized person
	authVersionOne := groupApiVersionOne.Group("/", AuthMiddleWare())
	authVersionOne.POST("users/add", AddVersionOne)

	// group version 2
	groupApiVersionTwo := router.Group("/v2")
	groupApiVersionTwo.GET("users", func(c *gin.Context) {
		c.JSON(http.StatusOK, "List Of group version two Users")
	})

	// User only can be added by authorized person
	authVersionTwo := groupApiVersionTwo.Group("/", AuthMiddleWare())
	authVersionTwo.POST("users/add", AddVersionTwo)

	// Listen and Server in
	_ = router.Run(":8081")
}

func AddVersionOne(c *gin.Context) {
	// AddUser Version 1
	c.JSON(http.StatusOK, "V1 User added")
}

func AddVersionTwo(c *gin.Context) {
	// AddUser Version 2
	c.JSON(http.StatusOK, "V2 User added")
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// here you can add your authentication method to authorize users.
		username := c.PostForm("user")
		password := c.PostForm("password")

		// checking the username and password
		if username == "foo" && password == "bar" {
			return
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
