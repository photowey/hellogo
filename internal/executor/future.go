package executor

import (
	"context"
)

const (
	single = 1
)

var _ Future = (*future)(nil)

// AwaitFunc {@code Future} {@code Await} func
type AwaitFunc func(ctx context.Context) (any, error)

// AwaitFuncFactory AwaitFunc 工厂
type AwaitFuncFactory func(ch chan any) AwaitFunc

// Future 异步编程模型接口
type Future interface {
	Await(ctxs ...context.Context) (any, error)
}

type future struct {
	await AwaitFunc
}

// Await 同步阻塞, 等待结果
func (f future) Await(ctxs ...context.Context) (any, error) {
	ctx := context.Background() // 默认的: ctx
	switch len(ctxs) {
	case single:
		ctx = ctxs[0]
	}

	return f.await(ctx)
}

// CreateAwaitFunc a func of {@code AwaitFuncFactory}
func CreateAwaitFunc(ch chan any, result *any) AwaitFunc {
	return func(ctx context.Context) (any, error) {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ch:
			return *result, nil
		}
	}
}

func NewFuture(ch chan any, result *any) Future {
	return &future{
		await: CreateAwaitFunc(ch, result),
	}
}
