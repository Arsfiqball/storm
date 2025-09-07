package main

import (
	"app/internal/system"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Arsfiqball/codec/talker"
	"github.com/spf13/viper"
)

// Example .config.yml
// ---
// # Server settings
// serve:
//   addr: ":3000"
//   monitor: ":8086"
// # Output format settings
// output: "json"
// # Database settings
// database:
//   url: "postgres://user:password@localhost:5432/app"
// # Redis settings
// redis:
//   url: "redis://localhost:6379/0"
//   consumer_group: "app-consumer-group"
//   consumer_prefix: "app-consumer"
//   # Supported Redis stream settings
//   max_idle_time: "1m"
//   block_time: "5s"
//   claim_interval: "30s"
//   claim_batch_size: 100
//   check_consumers_interval: "5m"
//   consumer_timeout: "10m"
//   nack_resend_sleep: "2s"
// # OpenTelemetry configuration
// telemetry:
//   service_name: "app-service"
//   zipkin_url: "" # Empty means use stdout exporter
//   sampling: "always" # Options: always, never, traceidratio, parentbased
//   sampling_ratio: 0.1 # Used when sampling is set to traceidratio
// ---

func main() {
	ctx := context.Background()

	// Use current directory for config
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search config in current directory with name ".config.yml" (without extension)
	viper.AddConfigPath(currentDir)
	viper.SetConfigName(".config.yml")

	// Read environment variables prefixed with APP_
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	monit := viper.GetString("serve.monitor")
	log.Printf("Monitor address for health checks: %s\n", monit)

	app, err := system.New(ctx)
	if err != nil {
		log.Fatalf("Error initializing application: %v", err)
	}

	proc := app.Serve()
	proc.MonitorAddr = monit

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	talker.Serve(proc, sig)
}
