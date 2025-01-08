package decorator

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

type loggingDecorator[Q, R any] struct {
	logger *logrus.Entry
	base   CommandHandler[Q, R]
}

func (l loggingDecorator[Q, R]) Handle(ctx context.Context, cmd Q) (result R, err error) {
	logger := l.logger.WithFields(logrus.Fields{
		"query":      generateActionName(cmd),
		"query_body": fmt.Sprintf("%#v", cmd),
	})

	logger.Debug("Executing")
	return l.base.Handle(ctx, cmd)
}
