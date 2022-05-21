package executor

import (
	"context"
	"errors"
)

var _ GoroutineExecutor = (*GoroutineExecutorPool)(nil)

type Task struct {
	runnable Runnable
	callable Callable
	resultCh chan any
	result   *any
	ctx      context.Context
}

func NewTaskr(task Runnable, ctx context.Context) Task {
	return Task{
		runnable: task, ctx: ctx,
	}
}

func NewTaskc(task Callable, ctx context.Context, ch chan any, result *any) (*Task, error) {
	if task == nil {
		return nil, errors.New("async runnable task queue is not enabled")
	}
	return &Task{
		callable: task, ctx: ctx, resultCh: ch, result: result,
	}, nil
}

type GoroutineExecutorPool struct {
	poolSize         int
	maxTaskQueueSize int
	asyncTaskQueue   chan Task
}

func NewGoroutineExecutorPool(poolSize, maxTaskQueueSize int) GoroutineExecutor {
	pool := &GoroutineExecutorPool{
		poolSize:         poolSize,
		maxTaskQueueSize: maxTaskQueueSize,
		asyncTaskQueue:   make(chan Task),
	}

	for i := 0; i < pool.poolSize; i++ {
		go func() {
			for {
				task, notClosed := <-pool.asyncTaskQueue
				if !notClosed {
					return
				} else {
					if task.runnable != nil {
						task.runnable(task.ctx)
					}
					if task.callable != nil {
						result := task.callable(task.ctx)
						if task.resultCh != nil {
							task.resultCh <- result
						}
					}
				}
			}
		}()
	}

	return pool
}

func (pool *GoroutineExecutorPool) Execute(task Runnable, ctx context.Context) error {
	if pool.asyncTaskQueue != nil {
		pool.asyncTaskQueue <- NewTaskr(task, ctx)
	} else {
		return errors.New("async runnable task queue is not enabled")
	}

	return nil
}

func (pool *GoroutineExecutorPool) Submit(task Callable, ctx context.Context) (Future, error) {
	if pool.asyncTaskQueue != nil {
		ch := make(chan any)
		var result any
		taskc, err := NewTaskc(task, ctx, ch, &result)
		if err != nil {
			return nil, err
		}
		pool.asyncTaskQueue <- *taskc
		return NewFuture(ch, &result), nil
	} else {
		return nil, errors.New("async callable task queue is not enabled")
	}
}
