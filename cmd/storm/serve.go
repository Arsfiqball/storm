package main

import (
	"app/internal/system"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Arsfiqball/csverse/talker"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Storm application server",
	Long:  `Start the Storm application server with the specified configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		// Get monitor address from viper
		monit := viper.GetString("serve.monitor")
		log.Printf("Monitor address for health checks: %s\n", monit)

		app, err := system.New(ctx)
		if err != nil {
			log.Fatalf("Error initializing Storm application: %v", err)
		}

		proc := app.Serve()
		proc.MonitorAddr = monit

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		talker.Serve(proc, sig)
	},
}

var serveHTTPCmd = &cobra.Command{
	Use:   "http",
	Short: "Start only the HTTP server component",
	Long:  `Start only the HTTP server component of the Storm application`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		// Get monitor address from viper
		monit := viper.GetString("serve.monitor")
		log.Printf("Monitor address for health checks: %s\n", monit)

		app, err := system.New(ctx)
		if err != nil {
			log.Fatalf("Error initializing Storm application: %v", err)
		}

		proc := app.ServeOnlyHTTP()
		proc.MonitorAddr = monit

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		talker.Serve(proc, sig)
	},
}

var serveListenerCmd = &cobra.Command{
	Use:   "listener",
	Short: "Start only the event listener component",
	Long:  `Start only the event listener component of the Storm application`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		// Get monitor address from viper
		monit := viper.GetString("serve.monitor")
		log.Printf("Monitor address for health checks: %s\n", monit)

		app, err := system.New(ctx)
		if err != nil {
			log.Fatalf("Error initializing Storm application: %v", err)
		}

		proc := app.ServeOnlyListener()
		proc.MonitorAddr = monit

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		talker.Serve(proc, sig)
	},
}

var serveWorkerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start only the worker component",
	Long:  `Start only the worker component of the Storm application`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		// Get monitor address from viper
		monit := viper.GetString("serve.monitor")
		log.Printf("Monitor address for health checks: %s\n", monit)

		app, err := system.New(ctx)
		if err != nil {
			log.Fatalf("Error initializing Storm application: %v", err)
		}

		proc := app.ServeOnlyWorker()
		proc.MonitorAddr = monit

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		talker.Serve(proc, sig)
	},
}

func init() {
	// Configure serve command flags
	serveCmd.Flags().StringP("addr", "a", ":3000", "Address to bind the server (default :3000)")
	serveCmd.Flags().StringP("monitor", "m", ":8086", "Monitor address for health checks")

	// Bind flags to viper
	viper.BindPFlag("serve.addr", serveCmd.Flags().Lookup("addr"))
	viper.BindPFlag("serve.monitor", serveCmd.Flags().Lookup("monitor"))

	// Set default values in viper
	viper.SetDefault("serve.addr", ":3000")
	viper.SetDefault("serve.monitor", ":8086")

	// Add subcommands to serve
	serveCmd.AddCommand(serveHTTPCmd)
	serveCmd.AddCommand(serveListenerCmd)
	serveCmd.AddCommand(serveWorkerCmd)
}
