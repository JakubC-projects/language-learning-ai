package chatgpt

type ChatCompletionRequest struct {
	Model            string    `json:"model,omitempty"`
	Messages         []Message `json:"messages,omitempty"`
	Stream           bool      `json:"stream,omitempty"`
	Temperature      int       `json:"temperature,omitempty"`
	TopP             int       `json:"top_p,omitempty"`
	MaxTokens        int       `json:"max_tokens,omitempty"`
	PresencePenalty  int       `json:"presence_penalty,omitempty"`
	FrequencyPenalty int       `json:"frequency_penalty,omitempty"`
}

type ChatCompletionResponse struct {
	Choices []Choice `json:"choices,omitempty"`
	Usage   Usage    `json:"usage,omitempty"`
}

type ChatCompletionChunkResponse struct {
	Choices []ChoiceChunk `json:"choices"`
}

type Choice struct {
	Index   int     `json:"index,omitempty"`
	Message Message `json:"message,omitempty"`
}

type ChoiceChunk struct {
	Index int     `json:"index,omitempty"`
	Delta Message `json:"delta,omitempty"`
}
type Usage struct {
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens,omitempty"`
}

type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}
