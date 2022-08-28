package fanout

import (
	"context"
	"errors"
	"runtime"
	"sync"

	"github.com/hellogo/pkg/logger"
)

const (
	DefaultNameFanout = "fanout"
	DefaultWorker     = 1 << 0
	DefaultBufferSize = 1 << 10
)

var (
	ErrQueueFull   = errors.New("fanout: chan full")
	ErrHandlerNull = errors.New("fanout: handle handler can't be bull")
)

type options struct {
	worker     uint
	bufferSize uint
}

func NewDefaultOptions() *options {
	return &options{
		worker:     DefaultWorker,
		bufferSize: DefaultBufferSize,
	}
}

type Option func(*options)

func WithWorker(worker uint) Option {
	return func(opt *options) {
		opt.worker = worker
	}
}

func WithBufferSize(bufferSize uint) Option {
	return func(opt *options) {
		opt.bufferSize = bufferSize
	}
}

type Task struct {
	handler func(ctx context.Context)
	ctx     context.Context
}

func NewTask(handler func(ctx context.Context), ctx context.Context) Task {
	return Task{
		handler: handler,
		ctx:     ctx,
	}
}

type Fanout struct {
	name    string
	queue   chan Task
	options *options
	waiter  sync.WaitGroup
	ctx     context.Context
	cancel  func()
}

func (queue *Fanout) Submit(ctx context.Context, handler func(ctx context.Context)) (err error) {
	if handler == nil {
		return ErrHandlerNull
	}

	if queue.ctx.Err() != nil {
		return queue.ctx.Err()
	}

	select {
	case queue.queue <- NewTask(handler, ctx):
	default:
		err = ErrQueueFull
	}

	return
}

func (queue *Fanout) Close() error {
	if err := queue.ctx.Err(); err != nil {
		return err
	}
	queue.cancel()
	queue.waiter.Wait()

	return nil
}

func NewFanout(name string, opts ...Option) *Fanout {
	if name == "" {
		name = DefaultNameFanout
	}

	optionz := NewDefaultOptions()
	for _, opt := range opts {
		opt(optionz)
	}

	queue := newFanout(name, optionz)
	queue.ctx, queue.cancel = context.WithCancel(context.Background())
	queue.waiter.Add(int(optionz.worker))

	for i := uint(0); i < optionz.worker; i++ {
		go func() {
			queue.newWorker()
		}()
	}

	return queue
}

func newFanout(name string, optionz *options) *Fanout {
	return &Fanout{
		name:    name,
		queue:   make(chan Task, optionz.bufferSize),
		options: optionz,
	}
}

func (queue *Fanout) newWorker() {
	defer queue.waiter.Done()
	for {
		select {
		case task := <-queue.queue:
			run := wrapFunc(task.handler)
			run(task.ctx)
		case <-queue.ctx.Done():
			return
		}
	}
}

func wrapFunc(fx func(ctx context.Context)) (runFunc func(context.Context)) {
	runFunc = func(ctx context.Context) {
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, 64*1024)
				buf = buf[:runtime.Stack(buf, false)]
				logger.Error("panic in fanout newWorker, err: %v, stack: %s", r, string(buf))
			}
		}()

		fx(ctx)
	}

	return
}
