package initcmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize eva configuration",
		Run: func(cmd *cobra.Command, args []string) {
			home, err := os.UserHomeDir()
			if err != nil {
				fmt.Println("❌ Could not find home directory:", err)
				os.Exit(1)
			}

			configDir := filepath.Join(home, ".eva")
			configFile := filepath.Join(configDir, "config.yaml")

			if _, err := os.Stat(configFile); err == nil {
				fmt.Println("⚠️  Config already exists at:", configFile)
				fmt.Print("Do you want to overwrite it? [y/N]: ")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				response := strings.ToLower(strings.TrimSpace(scanner.Text()))
				if response != "y" && response != "yes" {
					fmt.Println("❌ Init cancelled.")
					return
				}
			}

			scanner := bufio.NewScanner(os.Stdin)

			fmt.Print("🔑 Enter your Notion API Key: ")
			scanner.Scan()
			notionKey := strings.TrimSpace(scanner.Text())

			fmt.Print("🗂  Enter your Notion Database ID: ")
			scanner.Scan()
			notionDB := strings.TrimSpace(scanner.Text())

			if notionKey == "" || notionDB == "" {
				fmt.Println("❌ Both fields are required.")
				os.Exit(1)
			}

			os.MkdirAll(configDir, 0700)
			f, err := os.Create(configFile)
			if err != nil {
				fmt.Println("❌ Failed to create config file:", err)
				os.Exit(1)
			}
			defer f.Close()

			f.WriteString(fmt.Sprintf("notion_api_key: %s\nnotion_database_id: %s\n", notionKey, notionDB))

			fmt.Println("✅ Config created at:", configFile)
		},
	}
}
