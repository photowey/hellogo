package lifecycle

import (
	"context"
	"time"

	"github.com/hellogo/pkg/logger"
)

const defaultStopTimeout = 10 * time.Second

type Lifecycle interface {
	OnStop(fn func(ctx context.Context) error)
	Stop()
}

type lifecycle struct {
	onStop      []func(ctx context.Context) error
	stopTimeout time.Duration
}

func New() Lifecycle {
	return &lifecycle{
		onStop:      make([]func(ctx context.Context) error, 0),
		stopTimeout: defaultStopTimeout,
	}
}

func (lc *lifecycle) OnStop(fn func(ctx context.Context) error) {
	lc.onStop = append(lc.onStop, fn)
}

func (lc *lifecycle) Stop() {
	for _, fn := range lc.onStop {
		ctx, cancel := context.WithTimeout(context.Background(), lc.stopTimeout)

		err := fn(ctx)
		if err != nil {
			logger.Error("Failed to run cleanup func:%v", err)
		}

		cancel()
	}
}
