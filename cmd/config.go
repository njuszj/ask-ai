package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure API settings",
	Long: `Configure the API settings in the config file.
Settings will be saved to ~/.ask-ai/config.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get user's home directory
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error getting home directory: %v\n", err)
			os.Exit(1)
		}

		// Set up viper
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(filepath.Join(home, ".ask-ai"))

		// Read existing config
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				fmt.Println("Config file not found. Please run 'aa init' first.")
				os.Exit(1)
			}
			fmt.Printf("Error reading config: %v\n", err)
			os.Exit(1)
		}

		// Get flags
		apiKey, _ := cmd.Flags().GetString("api-key")
		apiEndpoint, _ := cmd.Flags().GetString("api-endpoint")
		temperature, _ := cmd.Flags().GetFloat64("temperature")

		// Update values if provided
		if apiKey != "" {
			viper.Set("api_key", apiKey)
		}
		if apiEndpoint != "" {
			viper.Set("api_endpoint", apiEndpoint)
		}
		if temperature != 0 {
			viper.Set("temperature", temperature)
		}

		// Save config
		if err := viper.WriteConfig(); err != nil {
			fmt.Printf("Error saving config: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Configuration saved successfully")
	},
}

func init() {
	configCmd.Flags().String("api-key", "", "set the api key")
	configCmd.Flags().String("api-endpoint", "", "set the api endpoint")
	configCmd.Flags().Float64("temperature", 0.7, "set the llm temperature")

	RootCmd.AddCommand(configCmd)
}
