package command

import (
	"context"
	slice "github.com/cglLaLaLa/SoVITS_web/common/genproto/slice_pb"
)

type CommandService interface {
	StartSlice(ctx context.Context, slice *slice.SliceRequest) (*slice.StartResponse, error)
}
