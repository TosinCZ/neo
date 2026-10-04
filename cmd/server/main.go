package main

import (
	"fmt"
	"github.com/Tosin/neo/internal/server"
)

func main() {
	s := server.New(":8080")

	if err := s.Start(); err != nil {
		fmt.Errorf("Failed start the server due to this error: %w\n", err)
	}
}