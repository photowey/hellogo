package async

import (
	"container/list"
	"context"
	"sync"
	"sync/atomic"
)

//
// @see * https://mp.weixin.qq.com/s/33uTlo3vIjFQAAokc6X5xw
//

const (
	defaultGoroutineMaxConcurrencyCount = ^uint32(0)
)

type GoroutinePoolOption func(*goroutinePoolOptions)

func WithGoroutinePoolMaxConcurrencyCount(c uint32) GoroutinePoolOption {
	return func(opts *goroutinePoolOptions) {
		opts.maxConcurrencyCount = c
	}
}

func WithGoroutinePoolWaitGroup(wg *sync.WaitGroup) GoroutinePoolOption {
	return func(opts *goroutinePoolOptions) {
		opts.wg = wg
	}
}

type goroutinePoolOptions struct {
	maxConcurrencyCount uint32
	wg                  *sync.WaitGroup
}

type GoroutinePool struct {
	ctx                      context.Context
	cancel                   context.CancelFunc
	goroutineChan            chan struct{}
	opts                     *goroutinePoolOptions
	jobsWaitQueue            list.List
	goroutineConcurrentCount uint32
	wg                       *sync.WaitGroup
}

func NewGoroutinePool(ctx context.Context, opts ...GoroutinePoolOption) *GoroutinePool {
	o := &goroutinePoolOptions{}
	for _, opt := range opts {
		opt(o)
	}

	initDefaultGoroutinePoolOptions(o)
	goCtx, cancel := context.WithCancel(ctx)
	return &GoroutinePool{
		ctx:           goCtx,
		cancel:        cancel,
		goroutineChan: make(chan struct{}, o.maxConcurrencyCount),
		opts:          o,
		wg:            o.wg,
	}
}

func initDefaultGoroutinePoolOptions(opts *goroutinePoolOptions) {
	if opts.maxConcurrencyCount == 0 {
		opts.maxConcurrencyCount = defaultGoroutineMaxConcurrencyCount
	}
}

func (gp *GoroutinePool) AddService(opts ...jobOption) *jobController {
	o := &jobOptions{}
	for _, opt := range opts {
		opt(o)
	}

	var ctx context.Context
	var cancel context.CancelFunc
	var j *job

	if o.serviceTimeout > 0 {
		ctx, cancel = context.WithTimeout(gp.ctx, o.serviceTimeout)
	} else {
		ctx, cancel = context.WithCancel(gp.ctx)
	}
	j = newJob(ctx, cancel, opts...)
	jc := &jobController{
		j: j,
	}
	if !o.shouldExecuteImmediately || gp.hasReachedConcurrencyLimit() {
		gp.jobsWaitQueue.PushBack(j)
		gp.tryExecuteJobQueue()
		return jc
	}
	gp.tryExecuteJob(j)
	return jc
}

func (gp *GoroutinePool) Wait() {
	gp.wg.Wait()
}

func (gp *GoroutinePool) Done() {
	<-gp.ctx.Done()
}

func (gp *GoroutinePool) tryExecuteJobQueue() {
	shouldExecuteJob := func() (*list.Element, bool) {
		if gp.hasReachedConcurrencyLimit() {
			return nil, false
		}
		jobElem := gp.jobsWaitQueue.Front()
		if jobElem == nil {
			return nil, false
		}

		return jobElem, true
	}

	for {
		jobElem, ok := shouldExecuteJob()
		if !ok {
			break
		}
		j, exists := jobElem.Value.(*job)
		if !exists {
			gp.jobsWaitQueue.Remove(jobElem)
			continue
		}
		gp.tryExecuteJob(j)
		if gp.hasReachedConcurrencyLimit() {
			gp.jobsWaitQueue.Remove(jobElem)
		}
	}
}

func (gp *GoroutinePool) tryExecuteJob(j *job) {
	var i uint32
	for i = 0; i < j.opts.serviceMaxConcurrencyCount; i++ {
		if gp.hasReachedConcurrencyLimit() {
			break
		}
		var s *BaseService
		var err error
		s, err = j.opts.serviceFactoryFunc(j.ctx, j.opts.args)
		if err != nil {
			return
		}
		s.j = j
		err = s.init()
		if err != nil {
			return
		}
		j.executeService(s, gp)
	}

	if j.opts.serviceMaxConcurrencyCount > i {
		gp.jobsWaitQueue.PushBack(j)
		return
	}
}

func (gp *GoroutinePool) incrementGoroutineConcurrentCount() {
	atomic.AddUint32(&gp.goroutineConcurrentCount, 1)
}

func (gp *GoroutinePool) decrementGoroutineConcurrencyCount() {
	atomic.AddUint32(&gp.goroutineConcurrentCount, ^uint32(0))
}

func (gp *GoroutinePool) GetGoroutineConcurrencyCount() uint32 {
	return atomic.LoadUint32(&gp.goroutineConcurrentCount)
}

func (gp *GoroutinePool) hasReachedConcurrencyLimit() bool {
	return gp.GetGoroutineConcurrencyCount() >= gp.opts.maxConcurrencyCount
}
