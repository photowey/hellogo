package executor

import (
	"context"
)

type GoroutineExecutor interface {
	Executor
	Submit(task Callable, ctx context.Context) (Future, error)
	// Submit(task Callable, ctx context.Context) error
}
