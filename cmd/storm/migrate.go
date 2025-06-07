package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Manage database migrations",
	Long:  `Run database migration operations like up, down, create, etc.`,
}

var migrateUpCmd = &cobra.Command{
	Use:   "up [steps]",
	Short: "Apply migrations",
	Long:  `Apply all or a limited number of up migrations`,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := getMigrator()
		if err != nil {
			fmt.Printf("Error creating migrator: %v\n", err)
			os.Exit(1)
		}

		if len(args) > 0 {
			steps, err := parseSteps(args[0])
			if err != nil {
				fmt.Printf("Error parsing steps: %v\n", err)
				os.Exit(1)
			}
			if err := m.Steps(steps); err != nil && !errors.Is(err, migrate.ErrNoChange) {
				fmt.Printf("Error applying %d migrations: %v\n", steps, err)
				os.Exit(1)
			}
			fmt.Printf("Applied %d migrations successfully\n", steps)
		} else {
			if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
				fmt.Printf("Error applying migrations: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Applied all migrations successfully")
		}
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down [steps]",
	Short: "Rollback migrations",
	Long:  `Rollback all or a limited number of migrations`,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := getMigrator()
		if err != nil {
			fmt.Printf("Error creating migrator: %v\n", err)
			os.Exit(1)
		}

		if len(args) > 0 {
			steps, err := parseSteps(args[0])
			if err != nil {
				fmt.Printf("Error parsing steps: %v\n", err)
				os.Exit(1)
			}
			if err := m.Steps(-steps); err != nil && !errors.Is(err, migrate.ErrNoChange) {
				fmt.Printf("Error rolling back %d migrations: %v\n", steps, err)
				os.Exit(1)
			}
			fmt.Printf("Rolled back %d migrations successfully\n", steps)
		} else {
			if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
				fmt.Printf("Error rolling back all migrations: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Rolled back all migrations successfully")
		}
	},
}

var migrateForceCmd = &cobra.Command{
	Use:   "force VERSION",
	Short: "Force migration version",
	Long:  `Force set specific migration version ignoring any errors`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m, err := getMigrator()
		if err != nil {
			fmt.Printf("Error creating migrator: %v\n", err)
			os.Exit(1)
		}

		version, err := parseVersion(args[0])
		if err != nil {
			fmt.Printf("Error parsing version: %v\n", err)
			os.Exit(1)
		}

		if err := m.Force(int(version)); err != nil {
			fmt.Printf("Error forcing migration version: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Forced migration version to %d successfully\n", version)
	},
}

var migrateVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get current migration version",
	Long:  `Print the current migration version of the database`,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := getMigrator()
		if err != nil {
			fmt.Printf("Error creating migrator: %v\n", err)
			os.Exit(1)
		}

		version, dirty, err := m.Version()
		if err != nil {
			fmt.Printf("Error getting migration version: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Current migration version: %d (dirty: %t)\n", version, dirty)
	},
}

var migrateCreateCmd = &cobra.Command{
	Use:   "create NAME",
	Short: "Create new migration files",
	Long:  `Create new UP and DOWN migration files with the given name`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		migrationsDir, _ := cmd.Flags().GetString("dir")
		ext, _ := cmd.Flags().GetString("ext")
		seq, _ := cmd.Flags().GetBool("seq")

		if migrationsDir == "" {
			migrationsDir = "database/postgresql/migrations"
		}

		// Ensure migrations directory exists
		if err := os.MkdirAll(migrationsDir, 0755); err != nil {
			fmt.Printf("Error creating migrations directory: %v\n", err)
			os.Exit(1)
		}

		// Generate filename timestamp or sequence
		var prefix string
		if seq {
			prefix = getNextSequence(migrationsDir, ext)
		} else {
			prefix = getTimestamp()
		}

		// Create migration files
		name := sanitizeName(args[0])
		upFilename := fmt.Sprintf("%s_%s.up.%s", prefix, name, ext)
		downFilename := fmt.Sprintf("%s_%s.down.%s", prefix, name, ext)

		upPath := fmt.Sprintf("%s/%s", migrationsDir, upFilename)
		downPath := fmt.Sprintf("%s/%s", migrationsDir, downFilename)

		// Create empty files
		if err := createEmptyFile(upPath); err != nil {
			fmt.Printf("Error creating up migration file: %v\n", err)
			os.Exit(1)
		}

		if err := createEmptyFile(downPath); err != nil {
			fmt.Printf("Error creating down migration file: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Created migration files:\n  %s\n  %s\n", upPath, downPath)
	},
}

func init() {
	// Configure migrate command flags
	migrateCmd.PersistentFlags().String("dir", "database/postgresql/migrations", "Directory containing migration files")
	migrateCmd.PersistentFlags().String("db", "", "Database connection string (defaults to config value)")

	// Create command specific flags
	migrateCreateCmd.Flags().String("ext", "sql", "File extension for migration files")
	migrateCreateCmd.Flags().Bool("seq", false, "Use sequential numbering instead of timestamps")

	// Add subcommands to migrate
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
	migrateCmd.AddCommand(migrateForceCmd)
	migrateCmd.AddCommand(migrateVersionCmd)
	migrateCmd.AddCommand(migrateCreateCmd)

	// Add to root command
	rootCmd.AddCommand(migrateCmd)
}

// Helper functions

func getMigrator() (*migrate.Migrate, error) {
	migrationsDir, _ := migrateCmd.PersistentFlags().GetString("dir")
	dbURL, _ := migrateCmd.PersistentFlags().GetString("db")

	if migrationsDir == "" {
		migrationsDir = "database/postgresql/migrations"
	}

	if dbURL == "" {
		dbURL = viper.GetString("database.url")
		if dbURL == "" {
			return nil, errors.New("database URL not provided and not found in configuration")
		}
	}

	// Create source URL
	sourceURL := fmt.Sprintf("file://%s", migrationsDir)

	// Create a new migrate instance
	m, err := migrate.New(sourceURL, dbURL)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func parseSteps(stepsStr string) (int, error) {
	var steps int
	_, err := fmt.Sscanf(stepsStr, "%d", &steps)
	if err != nil {
		return 0, errors.New("steps must be an integer")
	}
	if steps < 1 {
		return 0, errors.New("steps must be positive")
	}
	return steps, nil
}

func parseVersion(versionStr string) (uint, error) {
	var version uint
	_, err := fmt.Sscanf(versionStr, "%d", &version)
	if err != nil {
		return 0, errors.New("version must be a non-negative integer")
	}
	return version, nil
}

func getTimestamp() string {
	// Use current Unix timestamp
	return fmt.Sprintf("%d", time.Now().Unix())
}

func getNextSequence(dir string, ext string) string {
	// Read directory and find the highest sequence number
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "000001"
	}

	maxSeq := 0
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), "."+ext) {
			var seq int
			_, err := fmt.Sscanf(entry.Name(), "%d_", &seq)
			if err == nil && seq > maxSeq {
				maxSeq = seq
			}
		}
	}

	// Format with leading zeros
	return fmt.Sprintf("%06d", maxSeq+1)
}

func sanitizeName(name string) string {
	// Replace spaces and invalid chars with underscores
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "-", "_")
	return strings.ToLower(name)
}

func createEmptyFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
