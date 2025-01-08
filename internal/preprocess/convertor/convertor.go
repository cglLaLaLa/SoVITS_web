package convertor

import (
	sliceoapi "github.com/cglLaLaLa/SoVITS_web/common/client/slice"
	slice "github.com/cglLaLaLa/SoVITS_web/common/genproto/slice_pb"
	"github.com/cglLaLaLa/SoVITS_web/preprocess/entity"
)

type SliceStartConverter struct{}

func (s SliceStartConverter) EntityToProto(sliceStart *entity.SliceRequest) *slice.SliceRequest {
	return &slice.SliceRequest{
		InputPath:   sliceStart.InputPath,
		OutputRoot:  sliceStart.OutputRoot,
		Threshold:   sliceStart.Threshold,
		MinLength:   sliceStart.MinLength,
		MinInterval: sliceStart.MinInterval,
		HopSize:     sliceStart.HopSize,
		MaxSilKept:  sliceStart.MaxSilKept,
		XMax:        sliceStart.XMax,
		Alpha:       sliceStart.Alpha,
		NParts:      sliceStart.NParts,
	}
}

func (s SliceStartConverter) ClientToEntity(sliceStart *sliceoapi.SliceRequest) *entity.SliceRequest {
	return &entity.SliceRequest{
		InputPath:   sliceStart.SliceParams[0].InputPath,
		OutputRoot:  sliceStart.SliceParams[0].OutputRoot,
		Threshold:   sliceStart.SliceParams[0].Threshold,
		MinLength:   sliceStart.SliceParams[0].MinLength,
		MinInterval: sliceStart.SliceParams[0].MinInterval,
		HopSize:     sliceStart.SliceParams[0].HopSize,
		MaxSilKept:  sliceStart.SliceParams[0].MaxSilKept,
		XMax:        sliceStart.SliceParams[0].XMax,
		Alpha:       sliceStart.SliceParams[0].Alpha,
		NParts:      sliceStart.SliceParams[0].NParts,
	}
}
