package iterator

var _ Collection = (*collection)(nil)

type collection struct {
	index int
	array []any
}

func (u *collection) NewIterator() Iterator {
	return &iterator{
		array: u.array,
	}
}

func NewCollection(array []any, indexz ...int) Collection {
	index := 0
	switch len(indexz) {
	case 1:
		if indexz[0] > 0 && indexz[0] < len(array) {
			index = indexz[0]
		}
	}

	return &collection{
		index: index,
		array: array,
	}
}
