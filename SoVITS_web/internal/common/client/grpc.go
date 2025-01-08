package client

import (
	"context"
	"errors"
	slice "github.com/cglLaLaLa/SoVITS_web/common/genproto/slice_pb"
	logrus "github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"time"
)

const timeout = 10000000000 //默认10s超时
const addr = "127.0.0.1"    //测试阶段，先写本地地址，后期全放到配置文件里

func waitFor(addr string, timeout time.Duration) bool {
	portAvailable := make(chan struct{})
	timeOut := time.After(timeout)

	go func() {
		for {
			select {
			case <-timeOut:
				//return
			default:
			}
			_, err := net.Dial("tcp", addr)
			if err == nil {
				close(portAvailable)
				return
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()

	select {
	case <-portAvailable:
		return true
	case <-timeOut:
		return false
	}
}

func NewSliceGRPCClient(ctx context.Context) (client slice.SliceServiceClient, close func() error, err error) {
	if !waitFor(addr, timeout) {
		logrus.Error("addr:%s dial failed,timeout:%d", addr, timeout)
		return nil, nil, errors.New("grpc not available")
	}
	opts := grpcDialOpts(addr)
	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		return nil, func() error { return nil }, err
	}
	return slice.NewSliceServiceClient(conn), conn.Close, nil
}

func grpcDialOpts(_ string) []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithStatsHandler(otelgrpc.NewClientHandler())}
}
