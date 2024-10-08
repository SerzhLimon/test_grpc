package server

import (
	"context"

	"google.golang.org/grpc"

	"github.com/SerzhLimon/test_grpc/app/internal/cache"
	"github.com/SerzhLimon/test_grpc/app/internal/usecase"
	pb "github.com/SerzhLimon/test_grpc/app/test_grpc_proto"
)

type server struct {
	uc usecase.Usecase
}

func NewServer(uc *usecase.Usecase) *server {
	return &server{
		uc: *uc,
	}
}

func NewCore(cache cache.Rediscache) *grpc.Server {
	s := grpc.NewServer()
	uc := usecase.NewUsecase(cache)
	pb.RegisterPreviewServiceServer(s, NewServer(uc))
	return s
}

func (s *server) GetPreviewImage(ctx context.Context, req *pb.GetPreviewImageRequest) (*pb.GetPreviewImageResponse, error) {
	previewImage, err := s.uc.GetPreviewImage(req.GetUrl())
	if err != nil {
		return &pb.GetPreviewImageResponse{}, err
	}

	return &pb.GetPreviewImageResponse{Image: previewImage}, nil
}

func (s *server) GetPreviewImageSlice(
	ctx context.Context,
	req *pb.GetPreviewImageSliceRequest) (*pb.GetPreviewImageSliceResponse, error) {

	previewImages, err := s.uc.GetPreviewImageSlice(req.GetUrls())
	if err != nil {
		return &pb.GetPreviewImageSliceResponse{}, err
	}
	return &pb.GetPreviewImageSliceResponse{Images: previewImages}, nil
}
