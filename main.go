package main

import (
	//"os"
	//"os/signal"
	"encoding/json"
	"fmt"
	"log"
	"goWPC/controllers"
	"net/http"
	"github.com/gin-gonic/gin"
	"goWPC/frs"
	"github.com/sacOO7/gowebsocket"
	"github.com/gin-contrib/cors"
)

var listen = true

func initWS(socket gowebsocket.Socket,messages chan frs.FRSWSResponse){
//	interrupt := make(chan os.Signal, 1)
//	signal.Notify(interrupt, os.Interrupt)



	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Fatal("Received connect error - ", err)
	}

	socket.OnConnected = func(socket gowebsocket.Socket) {
		log.Println("Connected to server");
	}

	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {

		var wsResponse frs.FRSWSResponse
		json.Unmarshal([]byte(string(message)), &wsResponse)
		log.Println("Received message - "+wsResponse.VerifyFaceID + " " + wsResponse.PersonID )
		messages<-wsResponse
	}

	socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
		log.Println("Received ping - " + data)
	}

    socket.OnPongReceived = func(data string, socket gowebsocket.Socket) {
		log.Println("Received pong - " + data)
	}

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		log.Println("Disconnected from server \n")
		return
	}

	socket.Connect()

  //socket.SendText("Thoughtworks guys are awesome !!!!")
  /*
	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			socket.Close()
			return
		}
	}
	*/
		log.Println("\nExit ws routine \n")
}
func main() {

	router := gin.Default()
	router.Use(cors.Default())
  socket := gowebsocket.New("ws://172.22.20.175:80/fcsrecognizedresult")
	messages := make(chan frs.FRSWSResponse,10)
	go initWS(socket,messages)
	var frsClient = new(frs.FrsClient)
	frsClient.IP ="172.22.20.175:80";
	var sessionID = frsClient.FrsLogin("goapi","1qaz@WSX")
	fmt.Printf("FRS Login session ID= : %s\n", sessionID)
	v1 := router.Group("/api")
	{
		userController := new(controllers.UserController)
		userController.SessionID =sessionID;
		userController.Messages =messages;
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
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	router.Run(":9741")
	log.Println("Exit Program")

}
