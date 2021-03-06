package goroutine2

// ---------------------------------------------------------------- var of singleton factory by init

var factory Factory

// ---------------------------------------------------------------- init

func init() {
	factory = Factory{}
}

// ---------------------------------------------------------------- Factory

// IFactory 定义一个 Goroutine 的 工厂接口抽象
type IFactory interface {
	CreateGoroutine() Goroutine
}

// Factory {@code IFactory} 工厂核心实现
type Factory struct{}

// NewFactory 创建一个
func NewFactory() Factory {
	// 假的创建
	// 在初始化的时候 -> 已经创建好
	return factory
}

// CreateGoroutine 创建一个 {@code Goroutine}
func (factory Factory) CreateGoroutine(fx func(parameters ...any), options ...any) Goroutine {
	return NewGoroutine(fx, options...)
}

// ---------------------------------------------------------------- Goroutine

// IGoroutine 定义一个 Golang goroutine 抽象
type IGoroutine interface {
	Start()
	run()
}

// Goroutine 定义一个 Golang goroutine 抽象实现
type Goroutine struct {
	options []any                   // Goroutine 执行需要参数
	fx      func(parameters ...any) // 带可变参数执行: Goroutine
}

// NewGoroutine 创建 一个 Goroutine
func NewGoroutine(fx func(parameters ...any), options ...any) Goroutine {
	return Goroutine{
		fx:      fx,
		options: options,
	}
}

// Start 启动 一个 Goroutine
func (actor Goroutine) Start() Goroutine {
	actor.run()

	return actor
}

// Startpre 启动 一个 Goroutine, 并执行前置函数
func (actor Goroutine) Startpre(pre func()) Goroutine {
	actor.runAround(pre, func() {
	})

	return actor
}

// Startpost 启动 一个 Goroutine, 并执行后置函数
func (actor Goroutine) Startpost(post func()) Goroutine {
	actor.runAround(func() {
	}, post)

	return actor
}

// Startaround 启动 一个 Goroutine, 并执行环绕函数
func (actor Goroutine) Startaround(pre func(), post func()) Goroutine {
	actor.runAround(pre, post)

	return actor
}

// run 执行: Goroutine
// 将 run() 私有化 -> 不允许外界直接访问
func (actor Goroutine) run() {
	actor.runAround(func() {}, func() {
	})
}

// runAround 启动 一个 Goroutine, 并执行环绕函数 pre() | post()
func (actor Goroutine) runAround(pre func(), post func()) {
	go func(pre func(), post func(), parameters ...any) {
		pre()
		actor.fx(parameters...)
		post()
	}(pre, post, actor.options...)
}
