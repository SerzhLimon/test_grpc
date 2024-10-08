package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/SerzhLimon/test_grpc/app/internal/cache"
	"github.com/SerzhLimon/test_grpc/app/internal/server"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	cache := cache.NewRedisCache()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = cache.GetClient().Ping(ctx).Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Redis is running on :6379")

	s := server.NewCore(*cache)

	fmt.Println("Server is running on :8000")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
