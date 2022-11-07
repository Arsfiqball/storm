package main

import (
	"app/internal/system"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// File for liveness probe to watch
const LIVE_FILE = "/tmp/app_server_live"

func main() {
	serveCtx, cancelServeCtx := context.WithCancel(context.Background())

	log.Println("Initialize server...")

	// Create liveness file
	_, err := os.Create(LIVE_FILE)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize app
	app, err := system.New(serveCtx)
	if err != nil {
		log.Fatal(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig

		cleanCtx, cancelCleanCtx := context.WithTimeout(serveCtx, 30*time.Second)

		go func() {
			<-cleanCtx.Done()
			if cleanCtx.Err() == context.DeadlineExceeded {
				log.Fatal("Graceful shutdown timed out... Forcing exit now...")
			}
		}()

		log.Println("Gracefully shutdown server...")
		err := app.Clean(cleanCtx)
		if err != nil {
			log.Fatal(err)
		}

		// Remove liveness file
		err = os.Remove(LIVE_FILE)
		if err != nil {
			log.Fatal(err)
		}

		cancelCleanCtx()
		cancelServeCtx()
	}()

	log.Println("Starting server...")
	err = app.Serve(serveCtx)
	if err != nil {
		log.Fatal(err)
	}

	<-serveCtx.Done()
}
