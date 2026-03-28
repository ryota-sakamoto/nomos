package node

import (
	"context"
	"log"
	"net/http"
	"sync"
)

type Node struct {
	data *map[string]any
	mu   sync.Mutex
}

func (*Node) Run(ctx context.Context) error {
	server := http.Server{
		Addr: "12345",
	}

	go func() {
		log.Println("listen and serve")
		if err := server.ListenAndServe(); err != nil {
			return
		}
	}()

	<-ctx.Done()
	log.Println("shutting down")
	return server.Shutdown(ctx)
}
