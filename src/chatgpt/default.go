package chatgpt

import "github.com/JakubC-projects/language-learning-ai/src/config"

var Default = NewClient(config.C.ChatGptApiKey)
