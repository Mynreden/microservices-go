package server

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestNewGRPCServer(t *testing.T) {
	address := ":50051"
	grpcServer := NewGRPCServer(address)

	assert.NotNil(t, grpcServer)
	assert.Equal(t, address, grpcServer.address)
	assert.IsType(t, &grpc.Server{}, grpcServer.grpcServer)
}

func TestGRPCServer_Start(t *testing.T) {
	address := ":50052"
	grpcServer := NewGRPCServer(address)

	// Start the server in a separate goroutine
	go grpcServer.Start()

	// Allow some time for the server to start
	time.Sleep(100 * time.Millisecond)

	// Check if the server is listening
	conn, err := net.Dial("tcp", address)
	require.NoError(t, err)
	err = conn.Close()
	if err != nil {
		return
	}

	// Stop the server after the test
	defer grpcServer.Stop()
}

func TestGRPCServer_Stop(t *testing.T) {
	address := ":50053"
	grpcServer := NewGRPCServer(address)

	// Start the server in a separate goroutine
	go grpcServer.Start()

	// Allow some time for the server to start
	time.Sleep(100 * time.Millisecond)

	// Stop the server
	grpcServer.Stop()

	// Attempt to connect to the server
	_, err := net.Dial("tcp", address)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "refused")
}

func TestGRPCServer_GetServer(t *testing.T) {
	address := ":50054"
	grpcServer := NewGRPCServer(address)

	server := grpcServer.GetServer()

	assert.NotNil(t, server)
	assert.IsType(t, &grpc.Server{}, server)
}
