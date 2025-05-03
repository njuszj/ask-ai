package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/njuszj/ask-ai/internal/chat"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "aa [question]",
	Short: "A terminal tool to chat with LLM",
	Long: `ask-ai (aa) is an easy-use terminal tool to chat with llm.
You can use it in two modes:
1. Direct question mode: aa "your question here"
2. Interactive mode: aa -i`,
	Run: func(cmd *cobra.Command, args []string) {
		// Handle interactive mode
		if interactive, _ := cmd.Flags().GetBool("interactive"); interactive {
			chatCmd.Run(cmd, args)
			return
		}

		// Handle single question mode
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}

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

		question := strings.Join(args, " ")
		answer, err := chat.AskLLM(question, state)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(answer)
	},
}
