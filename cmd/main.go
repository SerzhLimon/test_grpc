package main

import (
	// "context"
	"fmt"
	"io"
	"net"
	"net/http"
	"regexp"

	"google.golang.org/grpc"

	pb "github.com/SerzhLimon/test_grpc/test_grpc_proto"
	"github.com/SerzhLimon/internal/server"
)

// type server struct {
// 	pb.UnimplementedPreviewServiceServer
// }

// func (s *server) GetPreviewImage(ctx context.Context, req *pb.PreviewRequest) (*pb.PreviewResponse, error) {
// 	videoURL := req.GetUrl()
// 	videoID, err := extractVideoID(videoURL)
// 	if err != nil {
// 		return nil, fmt.Errorf("invalid YouTube URL")
// 	}

// 	previewURL := fmt.Sprintf("https://img.youtube.com/vi/%s/0.jpg", videoID)
// 	previewImage, err := downloadImage(previewURL)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to download preview image")
// 	}

// 	return &pb.PreviewResponse{Image: previewImage}, nil
// }

func extractVideoID(videoURL string) (string, error) {
	re := regexp.MustCompile(`(?:https?:\/\/)?(?:www\.)?(?:youtube\.com\/(?:[^\/\n\s]+\/\S+\/|(?:v|e(?:mbed)?)\/|\S*?[?&]v=)|youtu\.be\/)([a-zA-Z0-9_-]{11})`)
	matches := re.FindStringSubmatch(videoURL)
	if len(matches) < 2 {
		return "", fmt.Errorf("invalid YouTube URL")
	}
	return matches[1], nil
}

func downloadImage(imageURL string) ([]byte, error) {
	resp, err := http.Get(imageURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download image: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}

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
