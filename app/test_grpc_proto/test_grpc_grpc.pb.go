// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: test_grpc.proto

package test

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PreviewService_GetPreviewImage_FullMethodName      = "/preview.PreviewService/GetPreviewImage"
	PreviewService_GetPreviewImageSlice_FullMethodName = "/preview.PreviewService/GetPreviewImageSlice"
)

// PreviewServiceClient is the client API for PreviewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PreviewServiceClient interface {
	GetPreviewImage(ctx context.Context, in *GetPreviewImageRequest, opts ...grpc.CallOption) (*GetPreviewImageResponse, error)
	GetPreviewImageSlice(ctx context.Context, in *GetPreviewImageSliceRequest, opts ...grpc.CallOption) (*GetPreviewImageSliceResponse, error)
}

type previewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPreviewServiceClient(cc grpc.ClientConnInterface) PreviewServiceClient {
	return &previewServiceClient{cc}
}

func (c *previewServiceClient) GetPreviewImage(ctx context.Context, in *GetPreviewImageRequest, opts ...grpc.CallOption) (*GetPreviewImageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPreviewImageResponse)
	err := c.cc.Invoke(ctx, PreviewService_GetPreviewImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *previewServiceClient) GetPreviewImageSlice(ctx context.Context, in *GetPreviewImageSliceRequest, opts ...grpc.CallOption) (*GetPreviewImageSliceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPreviewImageSliceResponse)
	err := c.cc.Invoke(ctx, PreviewService_GetPreviewImageSlice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PreviewServiceServer is the server API for PreviewService service.
// All implementations should embed UnimplementedPreviewServiceServer
// for forward compatibility.
type PreviewServiceServer interface {
	GetPreviewImage(context.Context, *GetPreviewImageRequest) (*GetPreviewImageResponse, error)
	GetPreviewImageSlice(context.Context, *GetPreviewImageSliceRequest) (*GetPreviewImageSliceResponse, error)
}

// UnimplementedPreviewServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPreviewServiceServer struct{}

func (UnimplementedPreviewServiceServer) GetPreviewImage(context.Context, *GetPreviewImageRequest) (*GetPreviewImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPreviewImage not implemented")
}
func (UnimplementedPreviewServiceServer) GetPreviewImageSlice(context.Context, *GetPreviewImageSliceRequest) (*GetPreviewImageSliceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPreviewImageSlice not implemented")
}
func (UnimplementedPreviewServiceServer) testEmbeddedByValue() {}

// UnsafePreviewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PreviewServiceServer will
// result in compilation errors.
type UnsafePreviewServiceServer interface {
	mustEmbedUnimplementedPreviewServiceServer()
}

func RegisterPreviewServiceServer(s grpc.ServiceRegistrar, srv PreviewServiceServer) {
	// If the following call pancis, it indicates UnimplementedPreviewServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PreviewService_ServiceDesc, srv)
}

func _PreviewService_GetPreviewImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPreviewImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PreviewServiceServer).GetPreviewImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PreviewService_GetPreviewImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PreviewServiceServer).GetPreviewImage(ctx, req.(*GetPreviewImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PreviewService_GetPreviewImageSlice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPreviewImageSliceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PreviewServiceServer).GetPreviewImageSlice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PreviewService_GetPreviewImageSlice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PreviewServiceServer).GetPreviewImageSlice(ctx, req.(*GetPreviewImageSliceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PreviewService_ServiceDesc is the grpc.ServiceDesc for PreviewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PreviewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "preview.PreviewService",
	HandlerType: (*PreviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPreviewImage",
			Handler:    _PreviewService_GetPreviewImage_Handler,
		},
		{
			MethodName: "GetPreviewImageSlice",
			Handler:    _PreviewService_GetPreviewImageSlice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test_grpc.proto",
}
