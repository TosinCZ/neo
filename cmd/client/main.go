package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		return fmt.Errorf("Failed to connect to the server due to this error: %w\n", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("ping"))
	if err != nil {
		return fmt.Errorf("Failed to write a message to the server due to this error:  %w\n", err):
	}

	response, err := io.ReadAll(conn)
	if err != nil {
		return fmt.Errorf("Failed to read the server due to this error:  %w\n", err):
	}

	fmt.Printf("Server responded with the following message: %s", response)
}