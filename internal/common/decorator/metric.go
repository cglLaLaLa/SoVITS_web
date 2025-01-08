package decorator

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type MetricClient interface {
	Inc(key string, value int)
}

type metricDecorator[Q, R any] struct {
	client MetricClient
	base   CommandHandler[Q, R]
}

func (m metricDecorator[Q, R]) Handle(ctx context.Context, cmd Q) (result R, err error) {
	start := time.Now()
	actionName := generateActionName(cmd)

	end := time.Since(start)
	m.client.Inc(actionName, int(end.Seconds()))

	return m.base.Handle(ctx, cmd)
}

func generateActionName(cmd any) string {
	return strings.Split(fmt.Sprintf("%T", cmd), ".")[1]
}
