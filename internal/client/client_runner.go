package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/SerzhLimon/test_grpc/test_grpc_proto"
)

const (
	directory = "../../images"
)

func main() {
	conn, err := grpc.NewClient(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("failed to connect grpc server", err)
	}
	defer conn.Close()

	client := pb.NewPreviewServiceClient(conn)

	async := makeFlags()

	if len(os.Args) < 2 || len(flag.Args()) == 0 {
		log.Fatalln("count of args must be > 0")
	}

	err = os.MkdirAll(directory, 0755)
	if err != nil {
		log.Fatalf("could not create directory: %v", err)
	}

	if async {
		req := &pb.GetPreviewImageSliceRequest{Urls: flag.Args()}
		resp, err := client.GetPreviewImageSlice(context.Background(), req)
		if err != nil {
			log.Fatalf("could not get preview image: %v", err)
		}
		err = saveImages(resp.GetImages(), directory)
	} else {
		req := &pb.GetPreviewImageRequest{Url: os.Args[1]}
		resp, err := client.GetPreviewImage(context.Background(), req)
		if err != nil {
			log.Fatalf("could not get preview image: %v", err)
		}
		filePath := filepath.Join(directory, fmt.Sprintf("preview.jpg"))
		fmt.Println(filePath)
		err = saveImage(resp.GetImage(), filePath)
	}
	
	fmt.Println("SUCCESS!\nImages saved in directory \"images\"")
}

func saveImage(image []byte, filePath string) error {
	return os.WriteFile(filePath, image, 0644)
}

func saveImages(images [][]byte, directory string) error {
	for i, image := range images {
		filePath := filepath.Join(directory, fmt.Sprintf("preview%d.jpg", i+1))
		if err := saveImage(image, filePath); err != nil {
			return err
		}
	}
	return nil
}

func makeFlags() bool {
	asyncFlag := flag.Bool("async", false, "")
	flag.Parse()
	return *asyncFlag
}
