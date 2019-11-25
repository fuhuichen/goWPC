package main

import (
	"goWPC/controllers"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api")
	{
		userController := new(controllers.UserController)
		v1.POST("/user/create", userController.Create)
		v1.POST("/user/delete", userController.Delete)
		v1.POST("/user/updateImage", userController.UpdateImage)

	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	router.Run(":9741")
}
