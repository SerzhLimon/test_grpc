package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/SerzhLimon/test_grpc/app/internal/cashe"
	"github.com/SerzhLimon/test_grpc/app/internal/server"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	cashe := cashe.NewRedisCache()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = cashe.GetClient().Ping(ctx).Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Redis is running on :6379")

	s := server.NewCore(*cashe)

	fmt.Println("Server is running on :8000")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
