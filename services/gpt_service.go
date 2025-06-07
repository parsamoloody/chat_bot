package services

import "exampe/chat/client"

func AskGPT(prompt string) (string, error) {
	return client.CallOpenAI(prompt)
}
