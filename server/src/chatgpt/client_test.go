package chatgpt

import (
	"fmt"
	"testing"

	"github.com/JakubC-projects/language-learning-ai/src/config"
	"github.com/stretchr/testify/assert"
)

func TestChatCompletions(t *testing.T) {
	c := NewClient(config.C.ChatGptApiKey)

	res, err := c.CreateChatCompletion(ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{Content: "Hello", Role: "user"},
		},
		Stream:    false,
		MaxTokens: 20,
	})
	assert.NoError(t, err)
	fmt.Println(res)
}

func TestChatCompletionsChan(t *testing.T) {
	c := NewClient(config.C.ChatGptApiKey)

	err := c.CreateChatCompletionChan(ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{Content: "Hello world", Role: "user"},
		},
		Stream:    false,
		MaxTokens: 100,
	}, make(chan ChatCompletionChunkResponse))
	assert.NoError(t, err)
}
