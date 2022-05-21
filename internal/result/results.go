package result

//
// 结果集设计
//

// ---------------------------------------------------------------- Result 设计

// Result 结果
//
// 成功 Ok
//
// 失败 error
//
// e.g.:
//
/*
func Request[R any](r R) Result[R] {
	return Result[R]{
		Ok:     r,
		Failed: nil,
	}
}
*/
type Result[T any] struct {
	Ok     T
	Failed error
}

// Expect 获取直接结果
//
// 如果：
//
// 1.结果集为成功,将返回 结果集包装的内容:
//
// 2.结果集为失败,将返回错误信息(Failed error)
func (r Result[T]) Expect(standBy T) (T, error) {
	if r.Failed != nil {
		return standBy, r.Failed
	}

	return r.Ok, nil
}
