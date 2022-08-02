package hicooflake

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

const (
	saltBit   = uint(1 << 3)        // 随机因子二进制位数
	saltShift = uint(1 << 3)        // 随机因子移位数
	incrShift = saltBit + saltShift // 自增数移位数
	randMax   = 1 << 8
)

const (
	defaultBase = uint64(1641063845000) // 2022-01-02 03:04:05
)

type Hicooflake struct {
	sync.Mutex
	incrBase uint64
	adam     uint8 // 随机因子一
	eve      uint8 // 随机因子二
}

func NewHicooflake(base ...uint64) *Hicooflake {
	incrBase := defaultBase
	switch len(base) {
	case 1:
		incrBase = base[0]
	}

	return &Hicooflake{incrBase: incrBase}
}

func NowHicooflake() *Hicooflake {
	now := uint64(time.Now().UnixNano() / 1e6)

	return NewHicooflake(now)
}

func (hf *Hicooflake) NextId() uint64 {
	hf.Lock()
	defer hf.Unlock()
	hf.incrBase++
	hf.adam = emit()
	hf.eve = emit()
	nextId := hf.incrBase<<incrShift | uint64(hf.adam<<saltShift) | uint64(hf.eve)

	return nextId
}

func (hf *Hicooflake) NextString() string {
	return fmt.Sprintf("%d", hf.NextId())
}

func emit() uint8 {
	seed, _ := rand.Int(rand.Reader, big.NewInt(randMax))

	return uint8(seed.Uint64())
}
