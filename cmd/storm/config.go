package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// configTemplate is the default config template
const configTemplate = `# Storm Framework Configuration

# Server settings
serve:
  addr: ":3000"
  monitor: ":8086"

# Output format settings
output: "json"

# Database settings
database:
  url: "postgres://user:password@localhost:5432/app"
`

var configCmd = &cobra.Command{
	Use:   "init-config",
	Short: "Initialize a default configuration file",
	Long:  `Create a default configuration file in the specified location or in the current directory`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("path")

		if configPath == "" {
			currentDir, err := os.Getwd()
			if err != nil {
				fmt.Println("Error getting current directory:", err)
				os.Exit(1)
			}
			configPath = filepath.Join(currentDir, ".storm.yaml")
		}

		// Check if file already exists
		if _, err := os.Stat(configPath); err == nil {
			overwrite, _ := cmd.Flags().GetBool("force")
			if !overwrite {
				fmt.Printf("Configuration file already exists at %s. Use --force to overwrite.\n", configPath)
				return
			}
		}

		// Write configuration template to file
		err := os.WriteFile(configPath, []byte(configTemplate), 0644)
		if err != nil {
			fmt.Println("Error writing configuration file:", err)
			os.Exit(1)
		}

		fmt.Printf("Configuration file created at %s\n", configPath)
	},
}

func init() {
	configCmd.Flags().String("path", "", "Path to create the configuration file (default is current directory/.storm.yaml)")
	configCmd.Flags().BoolP("force", "f", false, "Overwrite existing configuration file")

	// Add to root command
	rootCmd.AddCommand(configCmd)
}
