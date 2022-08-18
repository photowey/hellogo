package ringbuffer

type RingBuffer[T any] struct {
	next, prev *RingBuffer[T]
	Value      T
}

func (rb *RingBuffer[T]) init() *RingBuffer[T] {
	rb.next = rb
	rb.prev = rb

	return rb
}

func (rb *RingBuffer[T]) Next() *RingBuffer[T] {
	if rb.next == nil {
		return rb.init()
	}

	return rb.next
}

func (rb *RingBuffer[T]) Prev() *RingBuffer[T] {
	if rb.next == nil {
		return rb.init()
	}

	return rb.prev
}

func (rb *RingBuffer[T]) Move(n int) *RingBuffer[T] {
	if rb.next == nil {
		return rb.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			rb = rb.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			rb = rb.next
		}
	}

	return rb
}

func (rb *RingBuffer[T]) Link(s *RingBuffer[T]) *RingBuffer[T] {
	n := rb.Next()
	if s != nil {
		p := s.Prev()

		rb.next = s
		s.prev = rb
		n.prev = p
		p.next = n
	}

	return n
}

func (rb *RingBuffer[T]) Unlink(n int) *RingBuffer[T] {
	if n <= 0 {
		return nil
	}

	return rb.Link(rb.Move(n + 1))
}

func (rb *RingBuffer[T]) Len() int {
	n := 0
	if rb != nil {
		n = 1
		for p := rb.Next(); p != rb; p = p.next {
			n++
		}
	}

	return n
}

func (rb *RingBuffer[T]) Do(fx func(T)) {
	if rb != nil {
		fx(rb.Value)
		for p := rb.Next(); p != rb; p = p.next {
			fx(p.Value)
		}
	}
}

func NewRingBuffer[T any](n int) *RingBuffer[T] {
	if n <= 0 {
		return nil
	}
	r := new(RingBuffer[T])
	p := r
	for i := 1; i < n; i++ {
		p.next = &RingBuffer[T]{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p

	return r
}
