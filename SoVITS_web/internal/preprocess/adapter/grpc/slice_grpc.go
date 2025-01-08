package grpc

import (
	"context"
	slice "github.com/cglLaLaLa/SoVITS_web/common/genproto/slice_pb"
)

type SliceGRPC struct {
	client slice.SliceServiceClient
}

func NewSliceGRPC(client slice.SliceServiceClient) *SliceGRPC {
	return &SliceGRPC{client: client}
}

func (s SliceGRPC) StartSlice(ctx context.Context, slice *slice.SliceRequest) (*slice.StartResponse, error) {
	resp, err := s.client.StartSlice(ctx, slice)
	return resp, err
}
