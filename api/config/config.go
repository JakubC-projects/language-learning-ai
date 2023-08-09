package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ApiKey           string `json:"api_key,omitempty"`
	Model            string `json:"model,omitempty"`
	SystemPrompt     string `json:"system_prompt,omitempty"`
	Temperature      int    `json:"temperature,omitempty"`
	TopP             int    `json:"top_p,omitempty"`
	MaxTokens        int    `json:"max_tokens,omitempty"`
	PresencePenalty  int    `json:"presence_penalty,omitempty"`
	FrequencyPenalty int    `json:"frequency_penalty,omitempty"`
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