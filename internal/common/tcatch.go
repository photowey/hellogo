package common

import (
	"reflect"
)

type ExceptionHandler func(any)

type TryCatcher struct {
	catches map[reflect.Type]ExceptionHandler
	hold    func()
}

func Try(f func()) *TryCatcher {
	return &TryCatcher{
		catches: make(map[reflect.Type]ExceptionHandler),
		hold:    f,
	}
}

func (c *TryCatcher) Catch(err any, eh ExceptionHandler) *TryCatcher {
	c.catches[reflect.TypeOf(err)] = eh
	return c
}

func (c *TryCatcher) Finally(f func()) {
	defer func() {
		if err := recover(); err != nil {
			if h, ok := c.catches[reflect.TypeOf(err)]; ok {
				h(err)
			}
			f()
		}
	}()

	c.hold()
}
