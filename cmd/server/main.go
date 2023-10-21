package main

import (
	"app/internal/system"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Arsfiqball/talker/exco"
)

func main() {
	ctx := context.Background()

	app, err := system.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	proc := exco.Process{
		MonitorAddr: ":8086",
		Start:       app.Start,
		Live:        app.Live,
		Ready:       app.Ready,
		Stop:        app.Stop,
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	exco.Serve(proc, sig)
}
