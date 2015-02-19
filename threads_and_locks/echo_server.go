package main

import (
	"sync"
	"net"
	"regexp"
	"runtime"
)

// Notes
// - http://jan.newmarch.name/go/socket/chapter-socket.html#heading_toc_j_19
// - http://stackoverflow.com/questions/18405023/how-would-you-define-a-pool-of-goroutines-to-be-executed-at-once-in-golang?answertab=votes#tab-top

type ExecutorService struct {
	MaxPoolSize int
	Jobs chan *ConnectionHandler
}

func NewExecutorService(maxPoolSize int) *ExecutorService {
	e := &ExecutorService{
		MaxPoolSize: maxPoolSize,
		Jobs: make(chan *ConnectionHandler),
	}
	for i:=0; i<e.MaxPoolSize; i++ {
		go Worker(e.Jobs)
	}
	return e
}

func Worker(jobs chan *ConnectionHandler) {
	for job := range jobs {
		job.Run()
	}
}

func (e *ExecutorService) execute(handler *ConnectionHandler) {
	e.Jobs <- handler
}

type ConnectionHandler struct {
	Conn net.Conn
	Pool *sync.WaitGroup
}

// ConnectionHandler constructor
func NewConnectionHandler(conn net.Conn) *ConnectionHandler {
	return &ConnectionHandler{Conn: conn}
}

func (h *ConnectionHandler) Run() {
	// On exit of function
	// - Mark thread as done
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

	poolSize := runtime.NumCPU() * 2

	executor := NewExecutorService(poolSize)

	tcpAddr, _ := net.ResolveTCPAddr("tcp4", "localhost:4567")
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	
	for {
		conn, _ := listener.Accept()
		executor.execute(NewConnectionHandler(conn))
	}
}
