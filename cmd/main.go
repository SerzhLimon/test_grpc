package main

import (
	// "context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	pb "github.com/SerzhLimon/test_grpc/test_grpc_proto"
	// serv "github.com/SerzhLimon/test_grpc/internal/server"
)




func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterPreviewServiceServer(s, &server{})

	fmt.Println("Server is running on :8080")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
