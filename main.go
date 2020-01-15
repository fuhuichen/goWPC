package main

import (
	//"os"
	//"os/signal"
	"log"
	"goWPC/controllers"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

var listen = true





func main() {

	router := gin.Default()
	router.Use(cors.Default())
	/*
  socket = gowebsocket.New("ws://172.22.20.175:80/fcsrecognizedresult")
	messages := make(chan frs.FRSWSResponse,10)
	go initWS(socket,messages)
	var frsClient = new(frs.FrsClient)
	frsClient.IP ="172.22.20.175:80";
	var sessionID = frsClient.FrsLogin("goapi2","1qaz@WSX")
	fmt.Printf("FRS Login session ID= : %s\n", sessionID)
	*/
	v1 := router.Group("/api")
	{
		userController := new(controllers.UserController)
		userController.Init()
		v1.POST("/user/create", userController.Create)
		v1.POST("/user/delete", userController.Delete)
		v1.POST("/user/list", userController.List)
		v1.POST("/user/info", userController.Find)
		v1.POST("/user/face", userController.FindFace)
		v1.POST("/user/update", userController.Update)
		v1.POST("/user/updateCheck", userController.UpdateCheck)
		v1.POST("/user/updateImage", userController.UpdateImage)
		v1.POST("/fr/verification", userController.VerifyImage)
		v1.POST("/order/create", userController.CreateOrder)
		v1.POST("/order/list", userController.ListOrders)
		v1.POST("/order/last", userController.FindOrder)
		v1.POST("/order/isSpecialBonus", userController.FindBonus)
		v1.POST("/order/setSpecialBonus", userController.SetBonus)
		v1.POST("/product/list", userController.ListProducts)
		v1.POST("/product/update", userController.UpdateProduct)
		v1.POST("/restart", userController.Restart)
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	router.Run(":9741")
	log.Println("Exit Program")

}
