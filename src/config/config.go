package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ChatGptApiKey    string `envconfig:"CHAT_GPT_API_KEY"`
	Model            string `envconfig:"CHAT_GPT_MODEL"`
	SystemPrompt     string
	Temperature      float64
	TopP             float64
	MaxTokens        int
	PresencePenalty  float64
	FrequencyPenalty float64
	AdminPassword    string `envconfig:"ADMIN_PASSWORD"`
	ProjectID        string `envconfig:"PROJECT_ID"`
	Oauth            struct {
		ClientID     string `envconfig:"OAUTH_CLIENT_ID"`
		ClientSecret string `envconfig:"OAUTH_CLIENT_SECRET"`
		Domain       string `envconfig:"OAUTH_DOMAIN"`
	}
	Server struct {
		Host string `envconfig:"SERVER_HOST"`
	}
}

var C Config

func init() {
	config := Config{
		Temperature: 1,
		TopP:        1,
		MaxTokens:   1000,
	}

	envconfig.Process("", &config)

	C = config
}
