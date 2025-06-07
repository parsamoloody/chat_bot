package config

import (
	"log"
	"os"
)

var OpenAIKey string

func LoadEnv() {
	OpenAIKey = os.Getenv("OPENAI_API_KEY")
	if OpenAIKey == "" {
		log.Fatal("missing OPENAI_API_KEY")
	}
}