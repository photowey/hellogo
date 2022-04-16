package routine

type IGoroutine interface {
	Start()
	Run()
}

type Goroutine struct {
	Fx func()
}

func NewGoroutine(fx func()) Goroutine {
	return Goroutine{
		Fx: fx,
	}
}

func (actor Goroutine) Start() Goroutine {
	actor.Run()

	return actor
}

func (actor Goroutine) Run() {
	go func() {
		actor.Fx()
	}()
}
