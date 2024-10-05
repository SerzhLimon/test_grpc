package server

import (
	"context"

	"google.golang.org/grpc"

	usecase "github.com/SerzhLimon/test_grpc/internal/usecase"
	pb "github.com/SerzhLimon/test_grpc/test_grpc_proto"
)

type server struct {
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
