package main

import (
	"os"
	"path"

	"github.com/Arsfiqball/codec/access"
	"github.com/spf13/cobra"
)

var accessCmd = &cobra.Command{
	Use:   "access [module-name]",
	Short: "Generate an access module",
	Long:  `Generate a new access module with the specified name. The name must be alphanumeric and can include slashes for subdirectories.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		moduleName := args[0]

		rootDir, err := os.Getwd()
		if err != nil {
			cmd.PrintErrf("Error getting current working directory: %v\n", err)
			return
		}

		acc := access.Access{
			Name:    moduleName,
			RootDir: path.Join(rootDir, "pkg"),
		}

		if err := acc.Generate(); err != nil {
			cmd.PrintErrf("Error generating access module: %v\n", err)
			return
		}
	},
}

func init() {
	// Add to root command
	rootCmd.AddCommand(accessCmd)
}

