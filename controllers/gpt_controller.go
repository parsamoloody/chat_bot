package controllers

import (
	"exampe/chat/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PromptRequest struct {
	Prompt string `json:"prompt"`
}

func GPTHandler(c *gin.Context) {
	var req PromptRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	}

	answer, err := services.AskGPT(req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get answer"})
	}

	c.JSON(http.StatusOK, gin.H{"answer": answer})
}
