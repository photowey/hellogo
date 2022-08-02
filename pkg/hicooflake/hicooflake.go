package hicooflake

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

const (
	saltBit   = uint8(8)
	factorBit = uint8(8)
	incrShift = saltBit + factorBit
)

type Hicooflake struct {
	sync.Mutex
}

func NewHicooflake() *Hicooflake {
	return &Hicooflake{}
}

func (hf *Hicooflake) NextId() uint64 {
	hf.Lock()
	defer hf.Unlock()
	now := uint64(time.Now().UnixNano() / 1e6)
	salt := emit(1 << 8)
	factor := emit(1 << 8)
	nextId := now<<incrShift | uint64(salt<<(saltBit<<1)) | uint64(factor)

	return nextId
}

func (hf *Hicooflake) NextString() string {
	return fmt.Sprintf("%d", hf.NextId())
}

func emit(max int64) int64 {
	seed, _ := rand.Int(rand.Reader, big.NewInt(max))

	return seed.Int64()
}
