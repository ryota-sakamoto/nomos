package node_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ryota-sakamoto/nomos/internal/node"
)

func TestCreateMux(t *testing.T) {
	mux := node.CreateMux()
	ts := httptest.NewServer(mux)
	defer ts.Close()

	body := strings.NewReader("{}")
	res, err := http.Post(ts.URL+"/nomos.v1.NomosService/Healthz", "application/json", body)
	if err != nil {
		t.Fatalf("Failed to request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("Status code is not 200: %v", res.StatusCode)
	}
}
