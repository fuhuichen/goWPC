package main

import (
	"goWPC/controllers"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		user := new(controllers.UserController)
		v1.POST("/movies", movie.Create)

	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	router.Run()
}
