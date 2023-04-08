package async

type Service interface {
	Name() string
	init() error
	PreExecute() error
	Execute() error
	PostExecute() error
	Cleanup() error
}

type (
	ServiceInitFunc        func() error
	ServicePreExecute      func() error
	ServiceExecuteFunc     func() error
	ServicePostExecuteFunc func() error
	ServiceCleanupFunc     func() error
)

type ServiceOption func(*ServiceOptions)

func WithServiceInitFunc(f ServiceInitFunc) ServiceOption {
	return func(opts *ServiceOptions) {
		opts.initFunc = f
	}
}

func WithServicePreExecuteFunc(f ServicePreExecute) ServiceOption {
	return func(opts *ServiceOptions) {
		opts.preExecute = f
	}
}

func WithServiceExecuteFunc(f ServiceExecuteFunc) ServiceOption {
	return func(opts *ServiceOptions) {
		opts.executeFunc = f
	}
}

func WithServicePostExecuteFunc(f ServicePostExecuteFunc) ServiceOption {
	return func(opts *ServiceOptions) {
		opts.postExecute = f
	}
}

func WithServiceCleanupFunc(f ServiceCleanupFunc) ServiceOption {
	return func(opts *ServiceOptions) {
		opts.cleanupFunc = f
	}
}

type ServiceOptions struct {
	initFunc    ServiceInitFunc
	preExecute  ServicePreExecute
	executeFunc ServiceExecuteFunc
	postExecute ServicePostExecuteFunc
	cleanupFunc ServiceCleanupFunc
}
