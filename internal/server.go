package server

import (
	"fmt"
	"io"
	"net"
)

type Server struct {
	addr string
}

func New(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

/*
Responsible for establshing a listener on the tcp connection and handling the
connection when a client connects.
*/
func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("Listening on %s failed with the following error: %w\n", s.addr, err)
	}
	defer listener.Close()

	fmt.Printf("Server listening on %s\n", s.addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("When trying to accept a client the server failed with the following error: %v\n", err)
			continue
		}

		go s.handleConnection(conn) 
	}
}

/*
Responsible for reading the clients message and writing the response
*/
func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("A Client successfully connected on %s!\n", conn.RemoteAddr())

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		if err != io.EOF {
			fmt.Printf("Reading the clients message failed with the following error: %v\n", err)
		}
		return
	}

	message := string(buffer[:n])
	fmt.Printf("%q\n", message)

	_, err = conn.Write([]byte("pong\n"))
	if err != nil {
		fmt.Printf("Writing a response failed with the following error: %v\n", err)
		return
	}
}