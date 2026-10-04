package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
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
	listener, err := net.Listen("tcp", s.addr) // Establish a listener on the tcp connection

	ctx, stop := signal.NotifyContext( // Create a context that will be canceled when an interrupt signal is received
		context.Background(),
		os.Interrupt,
	)
	defer stop()

	if err != nil {
		return fmt.Errorf("Listening on %s failed with the following error: %w\n", s.addr, err)
	}
	defer listener.Close()

	fmt.Printf("Server listening on %s\n", s.addr)

	var wg sync.WaitGroup // Create a WaitGroup to keep track of active connections
	go func() {
		<-ctx.Done() // Wait for the context to be canceled 
		fmt.Println("Shutting down the server...")
		listener.Close() // Close the listener to stop accepting new connections
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			if ctx.Err() != nil {
				break // Exit the loop if the context is canceled (server is shutting down)
			}
			fmt.Printf("When trying to accept a client the server failed with the following error: %v\n", err)
			continue
		}

		wg.Add(1) // Increment the WaitGroup counter for a new connection
		go func() {
			defer wg.Done() // Decrement the WaitGroup counter when the connection is done
			s.handleConnection(conn)
		}()
	}
	wg.Wait() // Wait for all active connections to finish
	fmt.Println("Server shut down.")
	return nil
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