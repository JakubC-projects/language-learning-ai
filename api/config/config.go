package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ApiKey           string  `json:"api_key,omitempty"`
	Model            string  `json:"model,omitempty"`
	SystemPrompt     string  `json:"system_prompt,omitempty"`
	Temperature      float64 `json:"temperature,omitempty"`
	TopP             float64 `json:"top_p,omitempty"`
	MaxTokens        int     `json:"max_tokens,omitempty"`
	PresencePenalty  float64 `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`
}

var C Config

func init() {
	config := Config{
		Temperature: 1,
		TopP:        1,
		MaxTokens:   1000,
	}
	data, err := os.ReadFile("./config.json")
	if err != nil {
		panic(fmt.Errorf("cannot open config: %w", err))
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(fmt.Errorf("cannot load config: %w", err))
	}
	C = config
}
