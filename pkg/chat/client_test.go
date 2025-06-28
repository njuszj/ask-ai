package chat_test

import (
	"testing"

	"github.com/njuszj/ask-ai/pkg/chat"
)

func TestCallOpenaiApi(t *testing.T) {
	// Create a new State object
	testCfg := chat.OpenaiApiConfig{
		Model: "gpt-4.1",
		Input: "What is the day after Monday?",
	}
	cli := chat.OpenaiClient{}
	response, err := cli.GetResponse(testCfg)
	if err != nil {
		t.Errorf("Error calling OpenAI API: %v", err)
	}
	t.Logf("Response: %s", response)
}
