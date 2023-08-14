package router

import (
	"io"

	"github.com/JakubC-projects/language-learning-ai/chatgpt"
	"github.com/JakubC-projects/language-learning-ai/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type MessageRequest struct {
	Messages []chatgpt.Message `json:"messages"`
}

var client = chatgpt.NewClient(config.C.ApiKey)

func LoadEndpoints(r gin.IRouter) {
	r.POST("/message", func(c *gin.Context) {
		var reqBody MessageRequest
		c.BindWith(&reqBody, binding.JSON)
		responseChan := make(chan chatgpt.ChatCompletionChunkResponse)

		var messages = reqBody.Messages

		if config.C.SystemPrompt != "" {
			systemMessage := chatgpt.Message{
				Role:    "system",
				Content: config.C.SystemPrompt,
			}
			messages = append([]chatgpt.Message{systemMessage}, messages...)
		}

		go client.CreateChatCompletionChan(chatgpt.ChatCompletionRequest{
			Model:            config.C.Model,
			Messages:         messages,
			MaxTokens:        config.C.MaxTokens,
			Temperature:      config.C.Temperature,
			TopP:             config.C.TopP,
			PresencePenalty:  config.C.PresencePenalty,
			FrequencyPenalty: config.C.FrequencyPenalty,
		}, responseChan)

		c.Stream(func(w io.Writer) bool {
			output, ok := <-responseChan
			if !ok {
				return false
			}
			if len(output.Choices) == 0 {
				return true
			}
			c.Writer.Write([]byte(output.Choices[0].Delta.Content))
			return true
		})
	})
}
