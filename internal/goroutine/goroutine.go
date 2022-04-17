package goroutine

// IGoroutine 定义一个 Golang goroutine 抽象
type IGoroutine interface {
	Start()
	Run()
}

// Goroutine 定义一个 Golang goroutine 抽象实现
type Goroutine struct {
	Fx      func()              // 非参数执行: Goroutine
	Fxp     func(params any)    // 带参数执行: Goroutine
	Fxo     func(params ...any) // 带可变参数执行: Goroutine -> 证明 可以用可变参数替代上面的: Fx Fxp
	Options []any
}

// NewGoroutine 创建 一个 Goroutine
func NewGoroutine(fx func(), options ...any) Goroutine {
	return Goroutine{
		Fx:      fx,
		Options: options,
	}
}

// NewGoroutinep 创建 一个 Goroutine
func NewGoroutinep(fxp func(params any), options ...any) Goroutine {
	return Goroutine{
		Fxp:     fxp,
		Options: options,
	}
}

// NewGoroutineo 创建 一个 带可变参数的 Goroutine
func NewGoroutineo(fxo func(params ...any), options ...any) Goroutine {
	return Goroutine{
		Fxo:     fxo,
		Options: options,
	}
}

// Start 启动 一个 Goroutine
func (actor Goroutine) Start() Goroutine {
	actor.Run()

	return actor
}

// Run 启动 一个 Goroutine
func (actor Goroutine) Run() {
	go func() {
		actor.Fx()
	}()
}

// Runp 带参数执行: Goroutine
func (actor Goroutine) Runp(params any) {
	go func(param any) {
		actor.Fxp(param)
	}(params)
}

// Runo 带可变参数执行: Goroutine
func (actor Goroutine) Runo(params ...any) {
	go func(param ...any) {
		actor.Fxo(param...)
	}(params)
}

// Runf 带可变参数执行: Goroutine
func (actor Goroutine) Runf() {
	go func(param ...any) {
		actor.Fxo(param...)
	}(actor.Options)
}
