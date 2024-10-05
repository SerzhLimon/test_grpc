package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/SerzhLimon/test_grpc/test_grpc_proto"
)

const (
	directory = "images"
)

func main() {
	conn, err := grpc.DialContext(context.Background(), ":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPreviewServiceClient(conn)
	fmt.Println("Input url:")
	var inputUrl string
	fmt.Scan(&inputUrl)
	req := &pb.GetPreviewImageRequest{Url: inputUrl}
	resp, err := client.GetPreviewImage(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get preview image: %v", err)
	}
	
	filename := "preview.jpg"
	filePath := filepath.Join(directory, filename)

	err = os.MkdirAll(directory, 0755)
	if err != nil {
		log.Fatalf("could not create directory: %v", err)
	}
	err = saveImage(resp.GetImage(), filePath)
	if err != nil {
		log.Fatalf("could not save image: %v", err)
	}

	fmt.Println("Image saved as preview.jpg")
}

func saveImage(image []byte, filePath string) error {
	return os.WriteFile(filePath, image, 0644)
}
