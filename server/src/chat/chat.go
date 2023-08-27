package chat

import "github.com/JakubC-projects/language-learning-ai/src/chatgpt"

var Chat = []chatgpt.Message{}

func Add(m chatgpt.Message) {
	Chat = append(Chat, m)
}
