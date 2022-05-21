package iterator

var _ Iterator = (*iterator)(nil)

type iterator struct {
	index int
	array []any
}

func (u *iterator) HasNext() bool {
	if u.index < len(u.array) {
		return true
	}

	return false
}

func (u *iterator) Next() any {
	if u.HasNext() {
		dst := u.array[u.index]
		u.index++

		return dst
	}

	return nil
}
