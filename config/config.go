package config

import (
	"log"
	"os"
)

var OpenAIKey string
var Port string

func LoadEnv() {
	OpenAIKey = os.Getenv("OPENAI_API_KEY")
	Port = os.Getenv("PORT")
	if OpenAIKey == "" {
		log.Fatal("missing OPENAI_API_KEY")
	}
	if Port == "" {
		log.Fatal("missing PORT")
	}
}