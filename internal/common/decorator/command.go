package decorator

import (
	"context"
	"github.com/sirupsen/logrus"
)

type CommandHandler[C any, Q any] interface {
	Handle(ctx context.Context, q C) (Q, error)
}

func ApplyCommandDecorators[H, R any](handler CommandHandler[H, R], logger *logrus.Entry, client MetricClient) CommandHandler[H, R] {
	return loggingDecorator[H, R]{
		logger: logger,
		base: metricDecorator[H, R]{
			client: client,
			base:   handler,
		},
	}
}
