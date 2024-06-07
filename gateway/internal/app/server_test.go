package app

import (
	"fmt"

	config "github.com/mynreden/microservices-go/gateway/internal/config"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	cfg := &config.Config{
		Addr: ":8080",
	}

	// Create a mock handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// Run the server in a separate goroutine
	go func() {
		err := Server(cfg, handler)
		assert.NoError(t, err)
	}()

	// Allow some time for the server to start
	time.Sleep(1 * time.Second)

	// Create a test request to the server
	resp, err := http.Get("http://localhost:8080")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()

	body := make([]byte, resp.ContentLength)
	resp.Body.Read(body)

	assert.Equal(t, "Hello, World!\n", string(body))
}
