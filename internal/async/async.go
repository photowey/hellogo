package async

import (
	"context"
)

const (
	single = 1
)

var _ Future = (*future)(nil)

type AwaitFunc func(ctx context.Context) any

func NewAwaitFunc(ch chan struct{}, result *any) AwaitFunc {
	return func(ctx context.Context) any {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ch:
			return *result
		}
	}
}

// Future 异步编程模型接口
type Future interface {
	Await(ctx ...context.Context) any
}

type future struct {
	await AwaitFunc
}

func (f future) Await(ctxs ...context.Context) any {
	ctx := context.Background()
	switch len(ctxs) {
	case single:
		ctx = ctxs[0] // 为什么这样设计? 应该在执行回调函数的时候 - 可能有隐式传参的需求
	}

	return f.await(ctx)
}

// Run executes the async function
func Run(fx func() any) Future {
	var result any
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		result = fx()
	}() // Goroutine pool?
	return future{
		await: NewAwaitFunc(ch, &result),
	}
}
