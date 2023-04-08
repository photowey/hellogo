package async

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type JobRestartPolicy int

const (
	JobRestartPolicyNever JobRestartPolicy = iota
	JobRestartPolicyAlways
	JobRestartPolicyOnFailure
)

type ServiceFactoryFunc func(context.Context, ...interface{}) (*BaseService, error)

type jobOption func(*jobOptions)

func WithServiceRestartPolicy(restartPolicy JobRestartPolicy) jobOption {
	return func(opts *jobOptions) {
		opts.restartPolicy = restartPolicy
	}
}

func WithServiceMaxConcurrencyCount(c uint32) jobOption {
	return func(opts *jobOptions) {
		opts.serviceMaxConcurrencyCount = c
	}
}

func WithServiceMaxPanicErrorCount(c uint32) jobOption {
	return func(opts *jobOptions) {
		opts.serviceMaxPanicErrorCount = c
	}
}

func WithServiceMaxErrorCount(c uint32) jobOption {
	return func(opts *jobOptions) {
		opts.serviceMaxErrorCount = c
	}
}

func WithServiceMaxRetriesCount(c uint32) jobOption {
	return func(opts *jobOptions) {
		opts.serviceMaxRetriesCount = c
	}
}

func WithServiceRestartDelay(d time.Duration) jobOption {
	return func(opts *jobOptions) {
		opts.serviceRestartDelay = d
	}
}

func WithServiceMaxRestartDelay(d time.Duration) jobOption {
	return func(opts *jobOptions) {
		opts.serviceMaxRestartDelay = d
	}
}

func WithServiceTimeout(d time.Duration) jobOption {
	return func(opts *jobOptions) {
		opts.serviceTimeout = d
	}
}

func WithServiceFactoryFunc(f ServiceFactoryFunc) jobOption {
	return func(opts *jobOptions) {
		opts.serviceFactoryFunc = f
	}
}

func WithServiceShouldExecuteImmediately(b bool) jobOption {
	return func(opts *jobOptions) {
		opts.shouldExecuteImmediately = b
	}
}

type jobOptions struct {
	restartPolicy              JobRestartPolicy
	serviceMaxConcurrencyCount uint32
	serviceMaxPanicErrorCount  uint32
	serviceMaxErrorCount       uint32
	serviceMaxRetriesCount     uint32
	serviceRestartDelay        time.Duration
	serviceMaxRestartDelay     time.Duration
	serviceTimeout             time.Duration
	serviceFactoryFunc         ServiceFactoryFunc
	args                       []interface{}
	shouldExecuteImmediately   bool
}

type job struct {
	sync.Mutex
	ctx                    context.Context
	cancel                 context.CancelFunc
	wg                     sync.WaitGroup
	done                   chan struct{}
	opts                   *jobOptions
	finishedServiceCount   uint32
	serviceRestartDelay    time.Duration
	servicePanicErrorCount uint32
	serviceErrorCount      uint32
	serviceRetriesCount    uint32
	err                    atomic.Value
	serviceConcurrentCount uint32
}

func newJob(ctx context.Context, cancel context.CancelFunc, opts ...jobOption) *job {
	o := &jobOptions{}

	for _, opt := range opts {
		opt(o)
	}

	j := &job{
		ctx:    ctx,
		cancel: cancel,
		opts:   o,
		done:   make(chan struct{}),
	}

	return j
}

func (j *job) executeService(s *BaseService, gp *GoroutinePool) {
	gp.goroutineChan <- struct{}{}
	j.wg.Add(1)
	gp.wg.Add(1)
	go func() {
		gp.incrementGoroutineConcurrentCount()
		defer func() {
			<-gp.goroutineChan
			gp.decrementGoroutineConcurrencyCount()
			defer gp.wg.Done()
			defer j.wg.Done()
		}()

		for {
			if err := j.safeRunService(s); err != nil {
				j.setError(err)
				j.incrementServiceErrorCount()
				if j.shouldRestartService() {
					j.incrementServiceRetriesCount()
					fmt.Printf("service：restart %d/%d\n",
						j.getServiceRetriesCount(), j.opts.serviceMaxRetriesCount)
					time.Sleep(j.opts.serviceRestartDelay)
				} else {
					break
				}
			} else {
				break
			}
		}
	}()
}

func (j *job) safeRunService(s *BaseService) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("service panic: %v", r)
			j.incrementServicePanicErrorCount()
		}
	}()
	if err = s.PreExecute(); err != nil {
		return
	} else if err = s.Execute(); err != nil {
		return
	} else if err = s.PostExecute(); err != nil {
		return
	}
	return
}

func (j *job) shouldRestartService() bool {
	switch j.opts.restartPolicy {
	case JobRestartPolicyNever:
		return false
	case JobRestartPolicyAlways:
		return true
	case JobRestartPolicyOnFailure:
		if j.opts.serviceMaxRetriesCount > 0 &&
			j.getServiceRetriesCount() >= j.opts.serviceMaxRetriesCount {
			return false
		} else if j.opts.serviceMaxPanicErrorCount > 0 &&
			j.getServicePanicErrorCount() >= j.opts.serviceMaxPanicErrorCount {
			return false
		} else if j.opts.serviceMaxErrorCount > 0 &&
			j.getServiceErrorCount() >= j.opts.serviceMaxErrorCount {
			return false
		}

		return true
	}
	return false
}

func (j *job) monitorJob() {
	select {
	case <-j.ctx.Done():
		j.cancel()
		j.wg.Wait()
		close(j.done)
	}
}

func (j *job) isDone() bool {
	select {
	case <-j.done:
		return true
	default:
		return false
	}
}

func (j *job) setError(err error) {
	j.err.Store(err)
}

func (j *job) error() error {
	if e, ok := j.err.Load().(error); ok {
		return e
	}
	return nil
}

func (j *job) incrementFinishedServiceCount() {
	atomic.AddUint32(&j.finishedServiceCount, 1)
}

func (j *job) getFinishedServiceCount() uint32 {
	return atomic.LoadUint32(&j.finishedServiceCount)
}

func (j *job) incrementServicePanicErrorCount() {
	atomic.AddUint32(&j.servicePanicErrorCount, 1)
}

func (j *job) getServicePanicErrorCount() uint32 {
	return atomic.LoadUint32(&j.servicePanicErrorCount)
}

func (j *job) incrementServiceErrorCount() {
	atomic.AddUint32(&j.serviceErrorCount, 1)
}

func (j *job) getServiceErrorCount() uint32 {
	return atomic.LoadUint32(&j.serviceErrorCount)
}

func (j *job) incrementServiceRetriesCount() {
	atomic.AddUint32(&j.serviceRetriesCount, 1)
}

func (j *job) getServiceRetriesCount() uint32 {
	return atomic.LoadUint32(&j.serviceRetriesCount)
}

func (j *job) incrementServiceConcurrencyCount() {
	atomic.AddUint32(&j.serviceConcurrentCount, 1)
}

func (j *job) getServiceConcurrencyCount() uint32 {
	return atomic.LoadUint32(&j.serviceConcurrentCount)
}

func (j *job) hasReachedServiceConcurrencyLimit() bool {
	return j.getServiceConcurrencyCount() >= j.opts.serviceMaxConcurrencyCount
}

type jobController struct {
	j *job
}

func (jc *jobController) Cancel() {
	jc.j.cancel()
}

func (jc *jobController) IsJobDone() bool {
	return jc.j.isDone()
}
