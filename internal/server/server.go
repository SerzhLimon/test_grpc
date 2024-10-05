package server

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	usecase "github.com/SerzhLimon/test_grpc/internal/usecase"
	pb "github.com/SerzhLimon/test_grpc/test_grpc_proto"
)

type server struct {
	pb.UnimplementedPreviewServiceServer
	uc usecase.Usecase
}

func NewServer(usecase usecase.Usecase) *server {
	return &server{
		uc: usecase,
	}
}

func NewCore() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterPreviewServiceServer(s, NewServer(usecase.Usecase{}))
	return s
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
