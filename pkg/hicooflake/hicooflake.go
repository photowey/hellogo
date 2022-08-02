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

const (
	defaultBase = uint64(1641063845000) // 2022-01-02 03:04:05
)

type Hicooflake struct {
	sync.Mutex
	incrBase uint64
}

func NewHicooflake(base ...uint64) *Hicooflake {
	incrBase := defaultBase
	switch len(base) {
	case 1:
		incrBase = base[0]
	}

	return &Hicooflake{
		incrBase: incrBase,
	}
}

func NowHicooflake() *Hicooflake {
	now := uint64(time.Now().UnixNano() / 1e6)

	return NewHicooflake(now)
}

func (hf *Hicooflake) NextId() uint64 {
	hf.Lock()
	defer hf.Unlock()
	hf.incrBase++
	salt := emit(1 << 8)
	factor := emit(1 << 8)
	nextId := hf.incrBase<<incrShift | uint64(salt<<(saltBit<<1)) | uint64(factor)

	return nextId
}

func (hf *Hicooflake) NextString() string {
	return fmt.Sprintf("%d", hf.NextId())
}

func emit(max int64) int64 {
	seed, _ := rand.Int(rand.Reader, big.NewInt(max))

	return seed.Int64()
}
