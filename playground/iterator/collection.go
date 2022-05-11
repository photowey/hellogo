package iterator

type Collection interface {
	NewIterator() Iterator
}
