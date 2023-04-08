package ringbuffer

import (
	"errors"
	"io"
	"sync"
)

//
// @see * https://mp.weixin.qq.com/s/wKQTjfCSN6tkT8453EtwKw
//

const (
	defaultMaxCapacity = 4096
	smallBufferSize    = 512
	maxInt             = int(^uint(0) >> 1)
	minRead            = 512
)

var (
	errInvalidBuffer = errors.New("invalid buffer")
	errInvalidWrite  = errors.New("invalid write result")
	errTooLarge      = errors.New("too large")
	errNegativeRead  = errors.New("reader returned negative count from read")
	errBufferFull    = errors.New("buffer full")
)

type Option func(array *RingBuffer)

func WithMaxBufferCapacity(c int) Option {
	return func(r *RingBuffer) {
		r.cap = c
	}
}

func WithAutoGrow(b bool) Option {
	return func(r *RingBuffer) {
		r.autoGrow = b
	}
}

func WithGrowth(n int) Option {
	return func(r *RingBuffer) {
		r.growth = n
	}
}

func WithUseMutex(b bool) Option {
	return func(r *RingBuffer) {
		r.useMutex = b
		r.mux = &sync.Mutex{}
	}
}

type RingBuffer struct {
	buf      []byte
	cap      int
	readPos  int
	writePos int
	empty    bool
	autoGrow bool
	growth   int
	mux      *sync.Mutex
	useMutex bool
}

func New(opts ...Option) *RingBuffer {
	r := &RingBuffer{
		empty: true,
	}

	for _, opt := range opts {
		opt(r)
	}

	if r.cap == 0 {
		r.cap = defaultMaxCapacity
	}

	r.buf = make([]byte, 0, r.cap)

	return r
}

func NewRingBuffer(buf []byte, opts ...Option) *RingBuffer {
	r := &RingBuffer{
		buf:   buf,
		empty: true,
	}

	for _, opt := range opts {
		opt(r)
	}

	if r.cap == 0 {
		r.cap = len(r.buf)
	}

	return r
}

func NewBufferString(s string, opts ...Option) *RingBuffer {
	r := &RingBuffer{
		buf:   []byte(s),
		empty: true,
	}

	for _, opt := range opts {
		opt(r)
	}

	if r.cap == 0 {
		r.cap = len(r.buf)
	}

	return r
}

func (r *RingBuffer) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, errInvalidBuffer
	}

	if r.empty {
		return 0, io.EOF
	}

	r.lock()
	defer r.unlock()

	size := len(p)

	if size > r.Size() {
		size = r.Size()
	}

	n := 0

	if r.writePos > r.readPos {
		n = copy(p, r.buf[r.readPos:r.readPos+size])
		r.readPos += n
	} else {
		readableSize := r.cap - r.readPos

		if size < readableSize {
			n = copy(p, r.buf[r.readPos:r.readPos+size])
		} else {
			n = copy(p, r.buf[r.readPos:r.cap])
			n += copy(p[readableSize:], r.buf[:size-readableSize])
		}
	}

	r.readPos = (r.readPos + size) % len(r.buf)

	if r.readPos == r.writePos {
		r.empty = true
	}

	if size > n {
		return n, io.ErrShortWrite
	} else if n > size {
		return n, errInvalidWrite
	}

	return n, nil
}

func (r *RingBuffer) ReadByte() (byte, error) {
	if r.empty {
		r.Reset()
		return 0, io.EOF
	}

	r.lock()
	defer r.unlock()

	c := r.buf[r.readPos]
	r.readPos = (r.readPos + 1) % len(r.buf)

	return c, nil
}

func (r *RingBuffer) ReadFrom(reader io.Reader) (n int64, err error) {
	r.lock()
	defer r.unlock()

	for {
		i := r.grow(minRead)
		r.buf = r.buf[:i]
		m, e := r.Read(r.buf[i:cap(r.buf)])

		if 0 > m {
			return n, errNegativeRead
		}

		r.empty = false

		r.buf = r.buf[:i+m]
		n += int64(m)

		if e == io.EOF {
			return n, nil
		}

		if e != nil {
			return n, e
		}
	}
}

func (r *RingBuffer) Write(p []byte) (n int, err error) {
	// 获得要写入数据的长度
	pLen := len(p)
	_, ok := r.tryGrowByReSlice(pLen)
	// 获得可用空间大小
	available := r.Available()

	if available == 0 && !r.autoGrow {
		return 0, errBufferFull
	}

	if !ok && r.autoGrow {
		r.grow(pLen)
		available = r.Available()
	}

	// 写入的数据大于可用空间大小
	if pLen > available {
		pLen = available
	}

	writePos := r.writePos

	if writePos > r.readPos || r.empty {
		n = copy(r.buf[writePos:], p[:pLen])
		// r.writePos += n
	} else {
		n = copy(r.buf[writePos:], p[:pLen])
	}

	r.writePos = (writePos + n) % r.cap
	r.empty = false

	if n != pLen {
		return n, io.ErrShortWrite
	}

	return
}

func (r *RingBuffer) WriteByte(c byte) error {
	if r.Available() == 0 && !r.autoGrow {
		return errBufferFull
	}

	m, ok := r.tryGrowByReSlice(1)

	if !ok {
		m = r.grow(1)
	}

	r.buf[m] = c
	return nil
}

func (r *RingBuffer) WriteString(s string) (n int, err error) {
	return r.Write([]byte(s))
}

func (r *RingBuffer) WriteTo(w io.Writer) (n int64, err error) {
	if nBytes := r.Size(); nBytes > 0 {
		m, e := w.Write(r.buf[r.readPos:])

		if m > nBytes {
			return int64(m), e
		}

		r.readPos += m
		n = int64(m)

		if e != nil {
			return n, e
		}

		if m != nBytes {
			return n, io.ErrShortWrite
		}
	}

	r.Reset()
	return n, nil
}

func (r *RingBuffer) Bytes() []byte {
	available := r.Available()

	if available == 0 {
		return nil
	}

	if r.writePos >= r.readPos {
		return r.buf[r.writePos:len(r.buf)]
	}

	return r.buf[r.writePos:r.readPos]
}

func (r *RingBuffer) Peek(n int) ([]byte, error) {
	if r.empty {
		return make([]byte, 0), nil
	}

	size := r.Size()

	if size == 0 {
		return make([]byte, 0), nil
	}

	if n > size {
		n = size
	}

	buf := make([]byte, n)

	if r.writePos > r.readPos {
		copy(buf, r.buf[r.readPos:r.readPos+n])
	} else {
		m := copy(buf, r.buf[r.readPos:])
		copy(buf[m:], r.buf[0:n-m])
	}

	return buf, nil
}

func (r *RingBuffer) Discard(n int) (int64, error) {
	if r.empty {
		return 0, nil
	}

	size := r.Size()

	if size == 0 {
		return 0, nil
	}

	if n > size {
		n = size
	}

	r.readPos = (r.readPos + n) % len(r.buf)

	if r.readPos == r.writePos {
		r.Reset()
	}

	return int64(n), nil
}

func (r *RingBuffer) Grow(n int) {
	if 0 >= n {
		return
	}

	m := r.grow(n)
	r.buf = r.buf[:m]
}

func (r *RingBuffer) Reset() {
	r.empty = true
	r.readPos = 0
	r.writePos = 0
	r.buf = r.buf[:0]
}

func (r *RingBuffer) Empty() bool {
	return r.empty
}

func (r *RingBuffer) Len() int {
	return len(r.buf) - r.readPos
}

func (r *RingBuffer) Cap() int {
	return r.cap
}

func (r *RingBuffer) Available() int {
	if r.empty {
		return r.cap
	}

	if r.cap > len(r.buf) {
		return (r.cap - len(r.buf)) + (len(r.buf) - r.writePos)
	}

	if r.writePos > r.readPos {
		return r.cap - r.writePos + r.readPos
	}

	return r.readPos - r.writePos
}

func (r *RingBuffer) Size() int {
	if (r.readPos == r.writePos) || r.empty {
		if r.empty {
			return 0
		}

		return len(r.buf)
	}

	if r.writePos > r.readPos {
		return r.writePos - r.readPos
	}

	return r.cap - r.readPos + r.writePos
}

func (r *RingBuffer) Full() bool {
	return r.Available() == 0
}

func (r *RingBuffer) lock() {
	if !r.useMutex {
		return
	}

	r.mux.Lock()
}

func (r *RingBuffer) unlock() {
	if !r.useMutex {
		return
	}

	r.mux.Unlock()
}

func (r *RingBuffer) tryGrowByReSlice(n int) (int, bool) {
	if l := len(r.buf); n <= cap(r.buf)-l {
		r.buf = r.buf[:l+n]
		return l, true
	}

	return 0, false
}

func (r *RingBuffer) grow(n int) int {
	if r.empty {
		r.Reset()
	}

	m, ok := r.tryGrowByReSlice(n)

	if ok {
		return m
	}

	if r.cap == 0 {
		r.cap = defaultMaxCapacity
	}

	if r.buf == nil && n < r.cap {
		r.buf = make([]byte, n, r.cap)
		return 0
	}

	if n <= (r.cap/2 - m) {
		copy(r.buf, r.buf[r.readPos:])
	} else if r.cap > (maxInt - r.cap - n) {
		panic(errTooLarge)
	} else {
		r.buf = r.growSlice(r.buf[r.readPos:], r.readPos+n)
	}

	r.readPos = 0
	r.buf = r.buf[:m+n]

	return m
}

func (r *RingBuffer) growSlice(b []byte, n int) []byte {
	defer func() {
		if recover() != nil {
			panic(errTooLarge)
		}
	}()

	c := len(b) + n

	if c < 2*cap(b) {
		c = 2 * cap(b)
	}

	buf := make([]byte, c)
	copy(buf, b)
	return buf[:len(b)]
}
