package command

import (
	"context"
	"errors"
	"github.com/cglLaLaLa/SoVITS_web/common/decorator"
	"github.com/cglLaLaLa/SoVITS_web/preprocess/convertor"
	"github.com/cglLaLaLa/SoVITS_web/preprocess/entity"
	"github.com/sirupsen/logrus"
)

type StartSlice struct {
	CustomerID string
	Request    *entity.SliceRequest
}

type SliceResult struct {
	SliceId string
}

type startSliceHandler struct {
	logger *logrus.Entry
	client CommandService
}

func (s startSliceHandler) Handle(ctx context.Context, q StartSlice) (*SliceResult, error) {
	resp, err := s.client.StartSlice(ctx, convertor.NewSliceStartConverter().EntityToProto(q.Request))
	if err == nil {
		s.logger.WithFields(logrus.Fields{
			"slice error": err,
		})
	}
	//todo 规定一下常量值
	if resp.Status == "1" {
		return &SliceResult{SliceId: q.CustomerID}, nil
	}
	return nil, errors.New("返回错误")
}

type StartSliceHandler decorator.CommandHandler[StartSlice, *SliceResult]

func NewStartSliceHandler(ctx context.Context, client CommandService, logger *logrus.Entry, metric decorator.MetricClient) StartSliceHandler {
	if client == nil {
		panic("[NewStartSliceHandler] client is nil")
	}
	return decorator.ApplyCommandDecorators[StartSlice, *SliceResult](
		startSliceHandler{client: client},
		logger,
		metric,
	)
}

func validate(ctx context.Context, q StartSlice) error {
	//todo
	return nil
}
