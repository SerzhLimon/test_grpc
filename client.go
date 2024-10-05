package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	pb "path/to/your/proto/package"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPreviewServiceClient(conn)
	req := &pb.PreviewRequest{Url: "https://www.youtube.com/watch?v=rloqQY9CT8I&t=532s"}
	resp, err := client.GetPreviewImage(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get preview image: %v", err)
	}

	// Save the image to a file or display it
	err = saveImage(resp.GetImage(), "preview.jpg")
	if err != nil {
		log.Fatalf("could not save image: %v", err)
	}

	fmt.Println("Image saved as preview.jpg")
}

func saveImage(image []byte, filename string) error {
	return ioutil.WriteFile(filename, image, 0644)
}
