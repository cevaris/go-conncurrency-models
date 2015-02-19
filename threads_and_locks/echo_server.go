package main

import (
	"sync"
	"net"
	"regexp"
)

// Notes
// - http://jan.newmarch.name/go/socket/chapter-socket.html#heading_toc_j_19
// - http://stackoverflow.com/questions/18405023/how-would-you-define-a-pool-of-goroutines-to-be-executed-at-once-in-golang?answertab=votes#tab-top

type ConnectionHandler struct {
	Conn net.Conn
	Pool *sync.WaitGroup
}

// ConnectionHandler constructor
func NewConnectionHandler(conn net.Conn, pool *sync.WaitGroup) *ConnectionHandler {
	return &ConnectionHandler{Conn: conn, Pool: pool}
}

func (h *ConnectionHandler) Run() {
	// On exit of function
	// - Mark thread as done
	// - Close connection to socket
	defer h.Conn.Close()
	defer h.Pool.Done()

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
	var wg sync.WaitGroup
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", "localhost:4567")
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	
	for {
		conn, _ := listener.Accept()
		wg.Add(1)
		go NewConnectionHandler(conn, &wg).Run()
		
	}
	wg.Wait()
}
