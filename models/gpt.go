package models

type ChatRequest struct {
	Model     string      `json:"model"`
	Messages   []ChatEntry `json:"messages"`
	MaxTokens int         `json:"max_tokens,omitempty"`
}

type ChatEntry struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Choices []struct {
		Message ChatEntry `json:"message"`
	} `json:"choices"`
}
