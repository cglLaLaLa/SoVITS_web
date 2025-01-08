package service

import (
	"context"
	"github.com/cglLaLaLa/SoVITS_web/common"
	"github.com/cglLaLaLa/SoVITS_web/common/client"
	"github.com/cglLaLaLa/SoVITS_web/preprocess/adapter/grpc"
	"github.com/cglLaLaLa/SoVITS_web/preprocess/app"
	"github.com/cglLaLaLa/SoVITS_web/preprocess/app/command"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) (app.Application, func()) {
	sliceClient, closeSliceClient, err := client.NewSliceGRPCClient(ctx)
	if err != nil {
		panic(err)
	}
	sliceGRPc := grpc.NewSliceGRPC(sliceClient)
	return newApplicationClient(ctx, sliceGRPc), func() {
		_ = closeSliceClient()
	}
}

func newApplicationClient(ctx context.Context, serviceClient command.CommandService) app.Application {
	logger := logrus.NewEntry(logrus.StandardLogger())
	metric := common.MyMetric{}
	return app.Application{
		Commands: app.Commands{
			Command: command.NewStartSliceHandler(ctx, serviceClient, logger, metric),
		},
	}
}
