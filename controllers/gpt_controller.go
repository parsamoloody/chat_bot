package controllers

import (
	"chat_bot/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PromptRequest struct {
	Prompt string `json:"prompt"`
}

func GPTHandler(c *gin.Context) {
	var req PromptRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON payload"})
		return
	}
	fmt.Println(req.Prompt)
	answer, err := services.AskGPT(req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get answer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": answer})
}
