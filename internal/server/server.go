package server

import (
	"context"
	"fmt"

	uc "github.com/SerzhLimon/test_grpc/internal/usecase"
	pb "github.com/SerzhLimon/test_grpc/test_grpc_proto"
)

type server struct {
	core pb.UnimplementedPreviewServiceServer
	uc   uc.Usecase
}

func NewServer(core pb.UnimplementedPreviewServiceServer, uc uc.Usecase) *server {
	return &server{
		core: core,
		uc:   uc.Usecase,
	}
}

func (s *server) GetPreviewImage(ctx context.Context, req *pb.GetPreviewImageRequest) (*pb.GetPreviewImageResponse, error) {
	videoURL := req.GetUrl()
	videoID, err := s.uc.ExtractVideoID(videoURL)
	if err != nil {
		return nil, fmt.Errorf("invalid YouTube URL")
	}

	previewURL := fmt.Sprintf("https://img.youtube.com/vi/%s/0.jpg", videoID)
	previewImage, err := s.uc.DownloadImage(previewURL)
	if err != nil {
		return nil, fmt.Errorf("failed to download preview image")
	}

	return &pb.GetPreviewImageResponse{Image: previewImage}, nil
}
