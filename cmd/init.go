package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the configuration file",
	Long: `Initialize the configuration file at ~/.ask-ai/config.yaml.
This command will create the config directory and file if they don't exist.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get user's home directory
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error getting home directory: %v\n", err)
			os.Exit(1)
		}

		// Create .ask-ai directory
		configDir := filepath.Join(home, ".ask-ai")
		if err := os.MkdirAll(configDir, 0755); err != nil {
			fmt.Printf("Error creating config directory: %v\n", err)
			os.Exit(1)
		}

		// Set up viper
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(configDir)

		// Set default values
		viper.SetDefault("temperature", 0.7)

		// Prompt for API endpoint
		var apiEndpoint string
		fmt.Print("Enter your API endpoint (e.g., https://api.openai.com/v1): ")
		fmt.Scanln(&apiEndpoint)
		viper.Set("api_endpoint", apiEndpoint)

		// Prompt for API key
		var apiKey string
		fmt.Print("Enter your API key: ")
		fmt.Scanln(&apiKey)
		viper.Set("api_key", apiKey)

		// Create config file if it doesn't exist
		configPath := filepath.Join(configDir, "config.yaml")
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			if err := viper.WriteConfigAs(configPath); err != nil {
				fmt.Printf("Error creating config file: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Created config file at %s\n", configPath)
		} else {
			if err := viper.WriteConfig(); err != nil {
				fmt.Printf("Error updating config file: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Updated config file at %s\n", configPath)
		}
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
