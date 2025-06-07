package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

var isolateCmd = &cobra.Command{
	Use:   "isolate",
	Short: "Manage isolated development environment",
	Long:  `Control Docker-based isolated development environment for testing and development`,
}

// Add expose flag variable
var exposePortsFlag bool

var isolateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Start the isolated environment",
	Long:  `Start the Docker-based isolated development environment`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting isolated environment...")

		// Build the docker compose command
		dockerArgs := []string{"compose", "-f", "docker-compose.isolated.yml"}

		// If expose flag is set, include the expose yml file
		if exposePortsFlag {
			dockerArgs = append(dockerArgs, "-f", "docker-compose.isolated.expose.yml")
			fmt.Println("Ports will be exposed to the host machine")
		}

		// Add the remaining arguments
		dockerArgs = append(dockerArgs, "up", "--build", "-d")

		dockerCmd := exec.Command("docker", dockerArgs...)
		dockerCmd.Stdout = os.Stdout
		dockerCmd.Stderr = os.Stderr

		if err := dockerCmd.Run(); err != nil {
			fmt.Printf("Error starting isolated environment: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Isolated environment started successfully")
	},
}

var isolateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop and remove the isolated environment",
	Long:  `Stop and remove the Docker-based isolated environment and its volumes`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stopping isolated environment...")

		dockerCmd := exec.Command("docker", "compose", "-f", "docker-compose.isolated.yml", "down", "--volumes")
		dockerCmd.Stdout = os.Stdout
		dockerCmd.Stderr = os.Stderr

		if err := dockerCmd.Run(); err != nil {
			fmt.Printf("Error stopping isolated environment: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Isolated environment stopped successfully")
	},
}

var isolateShellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Open a shell in the application container",
	Long:  `Start a shell session in the application container of the isolated environment`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Opening shell in application container...")

		dockerCmd := exec.Command("docker", "compose", "-f", "docker-compose.isolated.yml", "run", "--rm", "app", "sh")
		dockerCmd.Stdout = os.Stdout
		dockerCmd.Stderr = os.Stderr
		dockerCmd.Stdin = os.Stdin

		if err := dockerCmd.Run(); err != nil {
			fmt.Printf("Error opening shell: %v\n", err)
			os.Exit(1)
		}
	},
}

var isolateCheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check requirements for isolated environment",
	Long:  `Verify that necessary tools like Docker and Docker Compose are available on the system`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking requirements for isolated environment...")

		// Check if docker is available
		dockerCmd := exec.Command("docker", "--version")
		if err := dockerCmd.Run(); err != nil {
			fmt.Println("❌ Docker is not available. Please install Docker to use the isolated environment.")
			os.Exit(1)
		}
		fmt.Println("✅ Docker is available")

		// Check if docker compose is available
		composeCmd := exec.Command("docker", "compose", "version")
		if err := composeCmd.Run(); err != nil {
			fmt.Println("❌ Docker Compose is not available. Please install Docker Compose to use the isolated environment.")
			os.Exit(1)
		}
		fmt.Println("✅ Docker Compose is available")

		// Check if docker is running
		checkRunningCmd := exec.Command("docker", "info")
		if err := checkRunningCmd.Run(); err != nil {
			fmt.Println("❌ Docker daemon is not running. Please start Docker to use the isolated environment.")
			os.Exit(1)
		}
		fmt.Println("✅ Docker daemon is running")

		fmt.Println("All requirements for isolated environment are met! ✨")
	},
}

// Add project name flag variable for test command
var testProjectNameFlag string

var isolateTestCmd = &cobra.Command{
	Use:   "test",
	Short: "Run component tests in the isolated environment",
	Long:  `Start the isolated environment, run component tests, and shut down the environment when complete`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("=============== SETUP START =================")
		fmt.Println("Starting isolated environment for tests...")

		// Start the isolated environment
		dockerArgs := []string{"compose", "-f", "docker-compose.isolated.yml"}

		// If project name flag is set, add it to docker compose command
		if testProjectNameFlag != "" {
			dockerArgs = append(dockerArgs, "-p", testProjectNameFlag)
			fmt.Printf("Using project name: %s\n", testProjectNameFlag)
		}

		dockerArgs = append(dockerArgs, "up", "--build", "-d")

		dockerCmd := exec.Command("docker", dockerArgs...)
		dockerCmd.Stdout = os.Stdout
		dockerCmd.Stderr = os.Stderr

		if err := dockerCmd.Run(); err != nil {
			fmt.Printf("Error starting isolated environment: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("================ SETUP END ==================")

		fmt.Println("Let me take a break for 10 seconds...")
		time.Sleep(10 * time.Second)

		// Run component tests
		fmt.Println("================ TEST START =================")
		testArgs := []string{"compose", "-f", "docker-compose.isolated.yml"}

		// If project name flag is set, add it to docker compose command
		if testProjectNameFlag != "" {
			testArgs = append(testArgs, "-p", testProjectNameFlag)
		}

		testArgs = append(testArgs, "run", "--rm", "app", "go", "test", "./test/component/...", "-v", "-p", "1")

		testCmd := exec.Command("docker", testArgs...)
		testCmd.Stdout = os.Stdout
		testCmd.Stderr = os.Stderr

		testErr := testCmd.Run()
		fmt.Println("================= TEST END ==================")

		fmt.Println("Let me take a break for 2 seconds...")
		time.Sleep(2 * time.Second)

		// Shut down the isolated environment
		fmt.Println("============== SHUTDOWN START ===============")
		fmt.Println("Shutting down isolated environment...")
		downArgs := []string{"compose", "-f", "docker-compose.isolated.yml"}

		// If project name flag is set, add it to docker compose command
		if testProjectNameFlag != "" {
			downArgs = append(downArgs, "-p", testProjectNameFlag)
		}

		downArgs = append(downArgs, "down", "--volumes")

		downCmd := exec.Command("docker", downArgs...)
		downCmd.Stdout = os.Stdout
		downCmd.Stderr = os.Stderr

		if err := downCmd.Run(); err != nil {
			fmt.Printf("Error shutting down isolated environment: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("=============== SHUTDOWN END ================")

		// Exit with the same status as the test command
		if testErr != nil {
			fmt.Printf("Tests failed: %v\n", testErr)
			os.Exit(1)
		}

		fmt.Println("Component tests completed successfully")
	},
}

func init() {
	// Add expose flag to the up command
	isolateUpCmd.Flags().BoolVarP(&exposePortsFlag, "expose", "e", false, "Expose database and app ports to the host machine")

	// Add project name flag only to the test command
	isolateTestCmd.Flags().StringVarP(&testProjectNameFlag, "project-name", "p", "", "Specify the Docker Compose project name")

	// Add subcommands to isolate command
	isolateCmd.AddCommand(isolateUpCmd)
	isolateCmd.AddCommand(isolateDownCmd)
	isolateCmd.AddCommand(isolateShellCmd)
	isolateCmd.AddCommand(isolateCheckCmd)
	isolateCmd.AddCommand(isolateTestCmd) // Add the new test command

	// Add to root command
	rootCmd.AddCommand(isolateCmd)
}
