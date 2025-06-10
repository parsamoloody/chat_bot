package main

import (
	"chat_bot/config"
	"chat_bot/controllers"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	config.LoadEnv()

	r := gin.Default()
	
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	getRout := r.Group("api")
	getRout.POST("/chat", controllers.GPTHandler)

	r.Run("localhost:" + config.Port)
}