package main

import (
	"fmt"
	"net"

	"github.com/SerzhLimon/test_grpc/internal/server"
	// "github.com/SerzhLimon/test_grpc/internal/client/cashe"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	s := server.NewCore()

	fmt.Println("Server is running on :8000")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
