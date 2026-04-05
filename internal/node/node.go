package node

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	nomosv1 "github.com/ryota-sakamoto/nomos/gen/nomos/v1"
	"github.com/ryota-sakamoto/nomos/gen/nomos/v1/nomosv1connect"
)

type Node struct {
	data *map[string]any
	mu   sync.Mutex
}

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()

	validator := validate.NewInterceptor()

	path, handler := nomosv1connect.NewNomosServiceHandler(&Node{}, connect.WithInterceptors(validator))
	mux.Handle(path, handler)

	return mux
}

func Run(ctx context.Context) error {
	mux := CreateMux()
	server := &http.Server{
		Addr:    "localhost:12345",
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	go func() {
		log.Println("listen and serve")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("failed to run server: %v\n", err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down")
	return server.Shutdown(ctx)
}

func (n *Node) GetItem(ctx context.Context, req *nomosv1.GetItemRequest) (*nomosv1.GetItemResponse, error) {
	log.Println("receive GetItem", req)
	return &nomosv1.GetItemResponse{}, connect.NewError(connect.CodeNotFound, fmt.Errorf("item is not found"))
}

func (n *Node) PutItem(ctx context.Context, req *nomosv1.PutItemRequest) (*nomosv1.PutItemResponse, error) {
	log.Println("receive PutItem", req)
	return &nomosv1.PutItemResponse{}, nil
}

func (n *Node) Healthz(ctx context.Context, req *nomosv1.HealthzRequest) (*nomosv1.HealthzResponse, error) {
	return &nomosv1.HealthzResponse{}, nil
}
