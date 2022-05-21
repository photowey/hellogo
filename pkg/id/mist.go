package id

import (
	"crypto/rand"
	"math/big"
	"sync"
)

const (
	saltBit   = uint(8)             // 随机因子二进制位数
	saltShift = uint(8)             // 随机因子移位数
	incrShift = saltBit + saltShift // 自增数移位数
)

type Mist struct {
	sync.Mutex       // 互斥锁
	incr       int64 // 自增数
	saltA      int64 // 随机因子一
	saltB      int64 // 随机因子二
}

func NewMist() *Mist {
	mist := Mist{incr: 1}
	return &mist
}

func (mt *Mist) Generate() int64 {
	mt.Lock()
	defer mt.Unlock()
	mt.incr++
	randA, _ := rand.Int(rand.Reader, big.NewInt(255))
	mt.saltA = randA.Int64()
	randB, _ := rand.Int(rand.Reader, big.NewInt(255))
	mt.saltB = randB.Int64()
	mist := int64((mt.incr << incrShift) | (mt.saltA << saltShift) | mt.saltB)

	return mist
}
