package chat

type AiClientBasicComponents struct {
	ApiEndpoint string
	ApiKey      string
}

// OpenaiApiConfig holds the common config using to interact with Openai api
type OpenaiApiConfig struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

type DeepseekApiConfig struct {
	Input          string
	Prompt         string
	Model          string
	Temperature    float64
	MaxTokens      int64
	TopP           float64
	Timeout        int
	ReturnFullText bool
}
