package llm

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Config holds the configuration for the LLM client
type Config struct {
	APIKey      string
	APIEndpoint string
	ModelName   string
	Temperature float64
}

// Client represents an LLM client
type Client struct {
	config     Config
	httpClient *http.Client
}

// NewClient creates a new LLM client
func NewClient(config Config) *Client {
	return &Client{
		config: config,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Ask sends a question to the LLM and returns the response
func (c *Client) Ask(ctx context.Context, question string, opts ...Option) (string, error) {
	// Create a copy of the config
	config := c.config

	// Apply any provided options
	for _, opt := range opts {
		opt(&config)
	}

	// Validate config
	if config.APIKey == "" {
		return "", fmt.Errorf("API key is required")
	}

	if config.APIEndpoint == "" {
		return "", fmt.Errorf("API endpoint is required")
	}

	// Placeholder response
	return fmt.Sprintf("Response to: %s (Temperature: %.2f)", question, config.Temperature), nil
}

// Option is a function that modifies the Config
type Option func(*Config)

// WithTemperature sets the temperature for a single request
func WithTemperature(temp float64) Option {
	return func(c *Config) {
		c.Temperature = temp
	}
}

// WithAPIKey sets the API key for a single request
func WithAPIKey(key string) Option {
	return func(c *Config) {
		c.APIKey = key
	}
}

// WithAPIEndpoint sets the API endpoint for a single request
func WithAPIEndpoint(endpoint string) Option {
	return func(c *Config) {
		c.APIEndpoint = endpoint
	}
}
