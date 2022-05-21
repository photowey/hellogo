package executor

import (
	"context"
)

type Executor interface {
	Execute(task Runnable, ctx context.Context) error
}
