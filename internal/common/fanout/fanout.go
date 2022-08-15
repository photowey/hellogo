package fanout

import (
	"context"
	"runtime"
	"sync"

	"github.com/pkg/errors"
)

const (
	EmptyString       = ""
	DefaultNameFanout = "fanout"
)

const (
	K = 1024

	DefaultPoolSize        = 1 << 0
	DefaultQueueBufferSize = 1 << 10
	StackBufSize           = 1 << 6
)

var ErrFull = errors.New("fanout: buffer queue fulled")

type Option func(*options)

type options struct {
	poolSize      uint
	maxBufferSize uint
}

type task struct {
	ctx context.Context
	fx  func(ctx context.Context)
}

type Executor struct {
	ctx     context.Context // ctx
	name    string          // name
	queue   chan task       // queue
	options *options        // opts
	wg      sync.WaitGroup  //
	cancel  func()          // cancel func
}

func (pool *Executor) run() {
	defer pool.wg.Done()
	for {
		select {
		case task := <-pool.queue:
			wrap(task.fx)(task.ctx)
		case <-pool.ctx.Done():
			return
		}
	}
}

func wrap(fx func(ctx context.Context)) (wrap func(context.Context)) {
	wrap = func(ctx context.Context) {
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, StackBufSize*K)
				buf = buf[:runtime.Stack(buf, false)]
			}
		}()
		fx(ctx)
	}
	return
}

func (pool *Executor) Execute(ctx context.Context, fn func(ctx context.Context)) (err error) {
	if fn == nil || pool.ctx.Err() != nil {
		return pool.ctx.Err()
	}

	select {
	case pool.queue <- newTask(ctx, fn):
	default:
		err = ErrFull
	}

	return
}

func (pool *Executor) Close() error {
	if err := pool.ctx.Err(); err != nil {
		return err
	}
	pool.cancel()
	pool.wg.Wait()

	return nil
}

func newTask(ctx context.Context, fx func(c context.Context)) task {
	return task{
		ctx: ctx,
		fx:  fx,
	}
}

func New(name string, opts ...Option) *Executor {
	if name == EmptyString {
		name = DefaultNameFanout
	}
	optionz := newOptions(DefaultPoolSize, DefaultQueueBufferSize)
	for _, opt := range opts {
		opt(optionz)
	}
	executor := newExecutor(name, optionz.maxBufferSize, optionz)

	executor.ctx, executor.cancel = context.WithCancel(context.Background())
	executor.wg.Add(int(optionz.poolSize))

	for i := uint(0); i < optionz.poolSize; i++ {
		go executor.run()
	}

	return executor
}

func WithPoolSize(poolSize uint) Option {
	return func(otps *options) {
		otps.poolSize = poolSize
	}
}

func WithMaxBufferSize(maxBufferSize uint) Option {
	return func(otps *options) {
		otps.maxBufferSize = maxBufferSize
	}
}

func newOptions(poolSize uint, maxBufferSize uint) *options {
	return &options{
		poolSize:      poolSize,
		maxBufferSize: maxBufferSize,
	}
}

func newExecutor(name string, maxBufferSize uint, opts *options) *Executor {
	return &Executor{
		name:    name,
		queue:   make(chan task, maxBufferSize),
		options: opts,
	}
}
