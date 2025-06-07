package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "storm",
		Short: "Storm is an AI-friendly fullstack framework",
		Long: `Storm is an AI-friendly fullstack framework based on Go Programming language.
It provides a comprehensive suite of tools for building, deploying, and managing 
modern web applications with built-in AI capabilities.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to Storm Framework! Use --help for available commands.")
		},
	}
)

func init() {
	// Initialize config
	cobra.OnInitialize(initConfig)

	// Add config flag
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./.storm.yaml)")

	// Add subcommands here
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(serveCmd)

	// Add global flags here
	rootCmd.PersistentFlags().StringP("output", "o", "", "output format (json|yaml)")
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Use current directory for config
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in current directory with name ".storm" (without extension)
		viper.AddConfigPath(currentDir)
		viper.SetConfigName(".storm")
	}

	// Read environment variables prefixed with STORM_
	viper.SetEnvPrefix("STORM")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Storm Framework",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Storm Framework v0.1")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
