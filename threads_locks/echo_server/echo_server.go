package main

import (
	"net"
	"regexp"
	"runtime"
	"github.com/cevaris/go_concurrency_models/threads_locks/executor_service"
)

type ConnectionHandler struct {
	Conn net.Conn
}

// ConnectionHandler constructor
func NewConnectionHandler(conn net.Conn) *ConnectionHandler {
	return &ConnectionHandler{Conn: conn}
}

func (h *ConnectionHandler) Run() {
	// On exit of function
	// - Close connection to socket
	defer h.Conn.Close()
	
	var buffer [1024]byte
	for { // Run till recieve exit request
		// Read in chunk of data
		n, err := h.Conn.Read(buffer[0:]);
		if err != nil {
			break
		}
		// Write out chunk of data
		_, err2 := h.Conn.Write(buffer[0:n])
		if err2 != nil {
			break
		}
		// Capture "exit" request from client
		match, _ := regexp.MatchString("^exit", string(buffer[0:]))
		if match {
			break
		}
	}
}

func main() {
	// Create thread pool using ExecutorService-like interface
	poolSize := runtime.NumCPU() * 2
	executor := executor_service.NewExecutorService(poolSize)

	tcpAddr, _ := net.ResolveTCPAddr("tcp4", "localhost:4567")
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	
	for {
		conn, _ := listener.Accept()
		executor.Execute(NewConnectionHandler(conn))
	}
}
