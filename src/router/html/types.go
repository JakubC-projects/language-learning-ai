package html

import "github.com/JakubC-projects/language-learning-ai/src/chatgpt"

type Login struct {
	ErrorMessage string
}

type HomePage struct {
	Messages []chatgpt.Message
}
