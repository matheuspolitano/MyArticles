package main

import (
	"log"
	"time"
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
}

func (s *Server) Run() {
	log.Printf("Server running %s:%d", s.host, s.port)

}
func (s *Server) Stop() {
	log.Printf("Server has stopped %s:%d", s.host, s.port)
}

// NewLocalHost creates a new Server instance with optional port and timeout parameters.
// If port or timeout are not provided (nil), default values are used.
func NewLocalHost(port interface{}, timeout interface{}) *Server {
	defaultPort := 8080
	defaultTimeout := 3 * time.Second

	// Check and set port if provided
	actualPort := defaultPort
	if p, ok := port.(int); ok {
		actualPort = p
	}

	// Check and set timeout if provided
	actualTimeout := defaultTimeout
	if t, ok := timeout.(time.Duration); ok {
		actualTimeout = t
	}

	return &Server{
		host:    "127.0.0.1",
		port:    actualPort,
		timeout: actualTimeout,
	}
}

func main() {
	// Example usage of NewLocalHost without parameters, using default values
	localHostServer := NewLocalHost(9090, nil)
	localHostServer.Run()

	// After some operations, stop the server
	// localHostServer.Stop()
}
