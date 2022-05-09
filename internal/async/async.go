package async

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
type AwaitFuncFactory func(ch chan struct{}, result *any) AwaitFunc

// CreateAwaitFunc a func of {@code AwaitFuncFactory}
func CreateAwaitFunc(ch chan struct{}, result *any) AwaitFunc {
	return func(ctx context.Context) (any, error) {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ch:
			return *result, nil
		}
	}
}

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
		ctx = ctxs[0] // 为什么这样设计? 应该在执行回调函数的时候 - 可能有隐式传参的需求
	}

	return f.await(ctx)
}

// Run executes the async function
func Run(fx func() any) Future {
	return Runz(fx, CreateAwaitFunc)
}

// Runz executes the async function with custom {@code AwaitFunc} factory
func Runz(fx func() any, factory AwaitFuncFactory) Future {
	var result any
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		result = fx()
	}() // Goroutine pool?
	return future{
		await: factory(ch, &result),
	}
}
