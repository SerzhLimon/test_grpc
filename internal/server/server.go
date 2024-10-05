package server

import (
	"context"
	"fmt"

	pb "github.com/SerzhLimon/test_grpc/test_grpc_proto"
	uc "github.com/SerzhLimon/test_grpc/usecase"
)


type server struct {
	core pb.UnimplementedPreviewServiceServer
	uc 
}

func NewServer(core pb.UnimplementedPreviewServiceServer) *server {
	return &server{
		core: core,
	}
}

func (s *server) GetPreviewImage(ctx context.Context, req *pb.GetPreviewImageRequest) (*pb.GetPreviewImageResponse, error) {
	videoURL := req.GetUrl()
	videoID, err := extractVideoID(videoURL)
	if err != nil {
		return nil, fmt.Errorf("invalid YouTube URL")
	}

	previewURL := fmt.Sprintf("https://img.youtube.com/vi/%s/0.jpg", videoID)
	previewImage, err := downloadImage(previewURL)
	if err != nil {
		return nil, fmt.Errorf("failed to download preview image")
	}

	return &pb.PreviewResponse{Image: previewImage}, nil
}
