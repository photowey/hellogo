package iterator

import (
	`errors`

	`github.com/hellogo/internal/hashmap`
)

var _ Iterator = (*iterator)(nil)
var ErrIteratorClosed = errors.New("iterator channel closed")

type Iterator interface {
	Iterator() (hashmap.HashMap, error)
	Stop() error
}

type iterator struct {
	ch     chan hashmap.HashMap
	stopCh chan struct{}
}

func (iter *iterator) Iterator() (hashmap.HashMap, error) {
	select {
	case hm := <-iter.ch:
		return hm, nil
	case <-iter.stopCh:
		return nil, ErrIteratorClosed
	}
}

func (iter *iterator) Stop() error {
	select {
	case <-iter.stopCh:
	default:
		close(iter.stopCh)
	}

	return nil
}

func NewIterator(ch chan hashmap.HashMap, stopCh chan struct{}) Iterator {
	return &iterator{
		ch:     ch,
		stopCh: stopCh,
	}
}
