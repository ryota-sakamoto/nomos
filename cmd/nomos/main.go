package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/ryota-sakamoto/nomos/internal/node"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if err := node.Run(ctx); err != nil {
		panic(err)
	}

	log.Println("stopped")
}
