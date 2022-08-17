package decorate

import (
	"fmt"
	"reflect"
)

func Decorate(decoPtr, f interface{}) error {
	fn := reflect.ValueOf(f)
	decoratedFunc := reflect.ValueOf(decoPtr).Elem()
	logicFunc := func(in []reflect.Value) []reflect.Value {
		println("before do something...")
		rvt := fn.Call(in)
		println("after do something...")
		return rvt
	}
	rvt := reflect.MakeFunc(fn.Type(), logicFunc)
	decoratedFunc.Set(rvt)

	return nil
}

func f1(a, b, c int) (int, error) {
	return a + b + c, nil
}

func Run() {
	decorateF1 := f1
	_ = Decorate(&decorateF1, f1)
	ret, err := decorateF1(1, 2, 3)
	fmt.Println(ret, err)
}
