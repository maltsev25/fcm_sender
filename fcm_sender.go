package main

import (
	"fcm_sender/controllers"
	"fcm_sender/environment"
	"fcm_sender/servicecontainer"
	"github.com/gin-gonic/gin"
	"log"
)

var env = environment.Get()

func main() {
	if env.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/", controllers.About)
	router.GET("/heartbeat", controllers.HeartBeat)
	sender := servicecontainer.ServiceContainer().InjectSenderController()
	router.POST("/send", sender.SendNotification)

	log.Println("Server starting on localhost:" + env.PORT)

	if err := router.Run(":" + env.PORT); err != nil {
		log.Fatal(err)
	}
}
