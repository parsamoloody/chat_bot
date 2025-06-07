package main

import (
	"exampe/chat/config"
	"exampe/chat/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("/chat", controllers.GPTHandler)

	r.Run("localhost:3030")
}