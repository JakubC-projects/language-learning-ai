package chatgpt

type ChatCompletionRequest struct {
	Model            string    `json:"model,omitempty"`
	Messages         []Message `json:"messages,omitempty"`
	Stream           bool      `json:"stream,omitempty"`
	Temperature      float64   `json:"temperature,omitempty"`
	TopP             float64   `json:"top_p,omitempty"`
	MaxTokens        int       `json:"max_tokens,omitempty"`
	PresencePenalty  float64   `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64   `json:"frequency_penalty,omitempty"`
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
