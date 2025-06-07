package client

import (
	"bytes"
	"encoding/json"
	"exampe/chat/config"
	"exampe/chat/models"
	"fmt"
	"io"
	"net/http"
	"time"
)

func CallOpenAI(Prompt string) (string, error) {
	reqBody := models.ChatRequest{
		Model: "gpt-4o-mini",
		Messages: []models.ChatEntry{
			{Role: "user", Content: Prompt},
		},
		MaxTokens: 50,
	}
	data, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(data))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+config.OpenAIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 15 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return "", fmt.Errorf("OpenAI error: %s", body)
	}
	var response models.ChatResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return "", err
	}
	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no choices")
	}
	return response.Choices[0].Message.Content, nil
	
}
