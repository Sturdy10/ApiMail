package main

import (
	"mail/handlers"
	"mail/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "X-Auth-Token", "Authorization"}
	router.Use(cors.New(config))

	s := services.NewServiceAdapter()
	h := handlers.NewHanerhandlerAdapter(s)

	router.POST("/api/sendMail", h.MailChicCRMHandlers)
	err := router.Run(":8888")
	if err != nil {
		panic(err.Error())
	}
}
