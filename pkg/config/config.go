package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	APIKey      string  `mapstructure:"apiKey"`
	APIEndpoint string  `mapstructure:"apiEndpoint"`
	ModelName   string  `mapstructure:"modelName"`
	Temperature float64 `mapstructure:"temperature"`
}

func LoadConfig() (*Config, error) {
	// Set default values
	viper.SetDefault("temperature", 0.7)

	// Set config file name and paths
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Add config paths
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %w", err)
	}

	viper.AddConfigPath(filepath.Join(home, ".ask-ai"))
	viper.AddConfigPath(".")

	// Read environment variables
	viper.SetEnvPrefix("ASK_AI")
	viper.AutomaticEnv()
	viper.BindEnv("api_key")
	viper.BindEnv("api_endpoint")
	viper.BindEnv("temperature")

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

func SaveConfig(config *Config) error {
	viper.Set("api_key", config.APIKey)
	viper.Set("api_endpoint", config.APIEndpoint)
	viper.Set("temperature", config.Temperature)

	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}

	configDir := filepath.Join(home, ".ask-ai")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	configPath := filepath.Join(configDir, "config.yaml")
	if err := viper.WriteConfigAs(configPath); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
