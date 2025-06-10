package services

import "chat_bot/client"

func AskGPT(prompt string) (string, error) {
	return client.CallOpenAI(prompt)
}
