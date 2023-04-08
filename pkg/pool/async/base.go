package async

import (
	"context"
)

type BaseService struct {
	ctx    context.Context
	cancel context.CancelFunc
	opts   *ServiceOptions
	j      *job
}

func NewBaseService(ctx context.Context, opts ...ServiceOption) *BaseService {
	o := &ServiceOptions{}
	for _, opt := range opts {
		opt(o)
	}

	return &BaseService{
		ctx:  ctx,
		opts: o,
	}
}

func (bs *BaseService) Name() string {
	panic("always assign a name to a service")
}

func (bs *BaseService) init() error {
	if bs.j.opts.serviceTimeout > 0 {
		bs.ctx, bs.cancel = context.WithTimeout(bs.ctx, bs.j.opts.serviceTimeout)
	} else {
		bs.ctx, bs.cancel = context.WithCancel(bs.ctx)
	}

	if bs.opts.initFunc != nil {
		return bs.opts.initFunc()
	}
	return nil
}

func (bs *BaseService) PreExecute() error {
	if bs.opts.preExecute != nil {
		return bs.opts.preExecute()
	}
	return nil
}

//

func (bs *BaseService) Execute() error {
	if bs.opts.executeFunc != nil {
		return bs.opts.executeFunc()
	}
	return nil
}

func (bs *BaseService) PostExecute() error {
	if bs.opts.postExecute != nil {
		return bs.opts.postExecute()
	}
	return nil
}

func (bs *BaseService) Cleanup() error {
	if bs.opts.cleanupFunc != nil {
		return bs.opts.cleanupFunc()
	}
	return nil
}
