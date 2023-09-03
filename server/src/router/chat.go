package router

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/JakubC-projects/language-learning-ai/src/auth"
	"github.com/JakubC-projects/language-learning-ai/src/chatgpt"
	"github.com/JakubC-projects/language-learning-ai/src/config"
	"github.com/JakubC-projects/language-learning-ai/src/db"
	"github.com/gin-gonic/gin"
)

func GetChats(c *gin.Context) {
	authCtx := auth.GetUser(c)
	chats, err := db.GetChats(c.Request.Context(), authCtx.UserId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("cannot get chats: %w", err))
		return
	}
	c.JSON(http.StatusOK, chats)
}

type createChatRequest struct {
	Prompt string `json:"prompt"`
}

func createChat(c *gin.Context) {
	var reqBody createChatRequest
	err := c.BindJSON(&reqBody)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid request body: %w", err))
		return
	}

	messages := []chatgpt.Message{
		{Content: "", Role: "system"},
		{Content: reqBody.Prompt, Role: "system"},
	}

	responseChan := make(chan chatgpt.ChatCompletionChunkResponse)

	go chatgpt.Default.CreateTestChatCompletionChan(chatgpt.ChatCompletionRequest{
		Model:            config.C.Model,
		Messages:         messages,
		MaxTokens:        config.C.MaxTokens,
		Temperature:      config.C.Temperature,
		TopP:             config.C.TopP,
		PresencePenalty:  config.C.PresencePenalty,
		FrequencyPenalty: config.C.FrequencyPenalty,
	}, responseChan)

	response := chatgpt.Message{Role: "assistant"}

	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-responseChan; ok {
			response.Content += msg.Choices[0].Delta.Content
			fmt.Println("Send message", response)
			resJson, _ := json.Marshal(response)
			c.SSEvent("message", resJson)
			return true
		}
		c.SSEvent("close", nil)
		return false
	})
}
