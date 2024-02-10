package main

import (
	"app/internal/system"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Arsfiqball/csverse/talker"
)

func main() {
	ctx := context.Background()

	app, err := system.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	proc := talker.Process{
		MonitorAddr: ":8086",
		Start:       app.Start,
		Live:        app.Live,
		Ready:       app.Ready,
		Stop:        app.Stop,
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	talker.Serve(proc, sig)
}
