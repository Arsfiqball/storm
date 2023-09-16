package main

import (
	"app/internal/system"
	"context"
	"log"

	"github.com/Arsfiqball/talker/excode"
)

func main() {
	ctx := context.Background()

	app, err := system.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	excode.Run(ctx, app, excode.RunConfig{HealthCheckAddress: ":8086"})
}
