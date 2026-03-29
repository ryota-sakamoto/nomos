package node

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	nomosv1 "github.com/ryota-sakamoto/nomos/gen/nomos/v1"
	"github.com/ryota-sakamoto/nomos/gen/nomos/v1/nomosv1connect"
)

type Node struct {
	data *map[string]any
	mu   sync.Mutex
}

func Run(ctx context.Context) error {
	mux := http.NewServeMux()

	path, handler := nomosv1connect.NewNomosServiceHandler(&Node{})
	mux.Handle(path, handler)

	server := &http.Server{
		Addr:    "localhost:12345",
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	go func() {
		log.Println("listen and serve")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println("failed to run server: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down")
	return server.Shutdown(ctx)
}

func (*Node) Ping(ctx context.Context, req *nomosv1.PingRequest) (*nomosv1.PingResponse, error) {
	log.Println("receive ping", req.Name)
	return &nomosv1.PingResponse{
		Message: fmt.Sprintf("hi %s", req.Name),
	}, nil
}
