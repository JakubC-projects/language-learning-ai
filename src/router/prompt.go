package router

import (
	"bytes"
	"io"
	"net/http"

	"github.com/JakubC-projects/language-learning-ai/src/chat"
	"github.com/JakubC-projects/language-learning-ai/src/chatgpt"
	"github.com/JakubC-projects/language-learning-ai/src/config"
	"github.com/gin-gonic/gin"
)

type PromptRequest struct {
	Prompt string `form:"prompt"`
}

type MessageWithResponse struct {
	Message  chatgpt.Message
	Response chatgpt.Message
}

func loadMessages(r *gin.Engine) {
	r.POST("/message", func(c *gin.Context) {

		var reqBody PromptRequest
		c.Bind(&reqBody)

		msg := chatgpt.Message{Role: "user", Content: reqBody.Prompt}
		chat.Add(msg)

		c.HTML(http.StatusOK, "components/message-with-live-response.html", MessageWithResponse{Message: msg, Response: chatgpt.Message{Role: "assistant"}})
	})
	r.GET("/generate-response", func(c *gin.Context) {
		responseChan := make(chan chatgpt.ChatCompletionChunkResponse)

		go chatgpt.Default.CreateChatCompletionChan(chatgpt.ChatCompletionRequest{
			Model:            config.C.Model,
			Messages:         chat.Chat,
			MaxTokens:        config.C.MaxTokens,
			Temperature:      config.C.Temperature,
			TopP:             config.C.TopP,
			PresencePenalty:  config.C.PresencePenalty,
			FrequencyPenalty: config.C.FrequencyPenalty,
		}, responseChan)

		response := chatgpt.Message{Role: "assistant"}

		c.Stream(func(w io.Writer) bool {
			if msg, ok := <-responseChan; ok {
				writer := new(bytes.Buffer)
				response.Content += msg.Choices[0].Delta.Content

				htmlTemplate.ExecuteTemplate(writer, "components/message.html", response)
				c.SSEvent("message", writer.String())
				return true
			}

			writer := new(bytes.Buffer)
			htmlTemplate.ExecuteTemplate(writer, "components/message.html", response)
			c.SSEvent("close", writer.String())
			return false
		})
	})
}
