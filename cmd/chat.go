package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/njuszj/ask-ai/pkg/chat"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Start an interactive chat session",
	Long: `Start an interactive chat session with the LLM.
You can change settings during the session using commands:
- set temp <value>: Change temperature
- set key <value>: Change API key
- set endpoint <value>: Change API endpoint
- show: Show current settings
- exit: Exit the session`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load config
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error getting home directory: %v\n", err)
			os.Exit(1)
		}

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(filepath.Join(home, ".ask-ai"))

		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Config file not found. Please run 'aa init' first.")
			os.Exit(1)
		}

		// Initialize chat state
		state := &chat.State{
			Temperature: viper.GetFloat64("temperature"),
			APIKey:      viper.GetString("api_key"),
			APIEndpoint: viper.GetString("api_endpoint"),
		}

		fmt.Println("Starting interactive chat...")
		fmt.Println("Type 'help' for available commands")

		for {
			fmt.Print("> ")
			var input string
			fmt.Scanln(&input)

			if input == "exit" {
				break
			}

			if input == "help" {
				chat.ShowHelp()
				continue
			}

			// Handle commands
			if strings.HasPrefix(input, ".set ") {
				chat.HandleSetCommand(input, state)
				continue
			}

			if input == "show" {
				chat.ShowSettings(state)
				continue
			}

			// Handle chat
			answer, err := chat.AskLLM(input, state)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}

			fmt.Println(answer)
		}
	},
}

func init() {
	RootCmd.AddCommand(chatCmd)
}
