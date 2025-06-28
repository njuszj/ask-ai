package chat

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/njuszj/ask-ai/pkg/llm"
)

// State holds the current state of the chat session
type State struct {
	Temperature float64
	APIKey      string
	APIEndpoint string
}

// ShowHelp displays available commands
func ShowHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  .set temp <value>     - Change temperature")
	fmt.Println("  .set key <value>      - Change API key")
	fmt.Println("  .set endpoint <value> - Change API endpoint")
	fmt.Println("  .show                 - Show current settings")
	fmt.Println("  .help                 - Show this help message")
	fmt.Println("  .exit                 - Exit the chat session")
}

// HandleSetCommand processes set commands
func HandleSetCommand(input string, state *State) {
	parts := strings.Split(input, " ")
	if len(parts) < 3 {
		fmt.Println("Invalid set command. Usage: set <param> <value>")
		return
	}

	param := parts[1]
	value := strings.Join(parts[2:], " ")

	switch param {
	case "temp":
		temp, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Printf("Invalid temperature value: %v\n", err)
			return
		}
		state.Temperature = temp
		fmt.Printf("Temperature set to %.2f\n", temp)
	case "key":
		state.APIKey = value
		fmt.Println("API key updated")
	case "endpoint":
		state.APIEndpoint = value
		fmt.Println("API endpoint updated")
	default:
		fmt.Println("Unknown parameter. Available parameters: temp, key, endpoint")
	}
}

// ShowSettings displays current settings
func ShowSettings(state *State) {
	fmt.Println("Current settings:")
	fmt.Printf("Temperature: %.2f\n", state.Temperature)
	if state.APIKey != "" {
		fmt.Println("API Key: [set]")
	} else {
		fmt.Println("API Key: [not set]")
	}
	if state.APIEndpoint != "" {
		fmt.Println("API Endpoint: [set]")
	} else {
		fmt.Println("API Endpoint: [not set]")
	}
}

// AskLLM sends a question to the LLM and returns the response
func AskLLM(question string, state *State) (string, error) {
	// Create LLM client
	client := llm.NewClient(llm.Config{
		APIKey:      state.APIKey,
		APIEndpoint: state.APIEndpoint,
		Temperature: state.Temperature,
	})

	// Make API call
	ctx := context.Background()
	response, err := client.Ask(ctx, question)
	if err != nil {
		return "", fmt.Errorf("failed to get response from LLM: %w", err)
	}

	return response, nil
}
