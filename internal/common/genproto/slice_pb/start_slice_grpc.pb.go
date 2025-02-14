// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: slice_pb/start_slice.proto

package slicepb

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
	SliceService_StartSlice_FullMethodName = "/slice.SliceService/StartSlice"
)

// SliceServiceClient is the client API for SliceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SliceServiceClient interface {
	StartSlice(ctx context.Context, in *SliceRequest, opts ...grpc.CallOption) (*StartResponse, error)
}

type sliceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSliceServiceClient(cc grpc.ClientConnInterface) SliceServiceClient {
	return &sliceServiceClient{cc}
}

func (c *sliceServiceClient) StartSlice(ctx context.Context, in *SliceRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, SliceService_StartSlice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SliceServiceServer is the server API for SliceService service.
// All implementations should embed UnimplementedSliceServiceServer
// for forward compatibility.
type SliceServiceServer interface {
	StartSlice(context.Context, *SliceRequest) (*StartResponse, error)
}

// UnimplementedSliceServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSliceServiceServer struct{}

func (UnimplementedSliceServiceServer) StartSlice(context.Context, *SliceRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartSlice not implemented")
}
func (UnimplementedSliceServiceServer) testEmbeddedByValue() {}

// UnsafeSliceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SliceServiceServer will
// result in compilation errors.
type UnsafeSliceServiceServer interface {
	mustEmbedUnimplementedSliceServiceServer()
}

func RegisterSliceServiceServer(s grpc.ServiceRegistrar, srv SliceServiceServer) {
	// If the following call pancis, it indicates UnimplementedSliceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SliceService_ServiceDesc, srv)
}

func _SliceService_StartSlice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SliceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SliceServiceServer).StartSlice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SliceService_StartSlice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SliceServiceServer).StartSlice(ctx, req.(*SliceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SliceService_ServiceDesc is the grpc.ServiceDesc for SliceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SliceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "slice.SliceService",
	HandlerType: (*SliceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartSlice",
			Handler:    _SliceService_StartSlice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slice_pb/start_slice.proto",
}
