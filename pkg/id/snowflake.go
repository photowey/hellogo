package id

import (
	"errors"
	"sync"
	"time"
)

const (
	twepoch        = int64(1483228800000)             // 开始时间截 (2017-01-01)
	workerIdBits   = uint(10)                         // 机器id所占的位数
	sequenceBits   = uint(12)                         // 序列所占的位数
	workerIdMax    = int64(-1 ^ (-1 << workerIdBits)) // 支持的最大机器id数量
	sequenceMask   = int64(-1 ^ (-1 << sequenceBits)) //
	workerIdShift  = sequenceBits                     // 机器id左移位数
	timestampShift = sequenceBits + workerIdBits      // 时间戳左移位数
)

type Snowflake struct {
	sync.Mutex
	timestamp int64
	workerId  int64
	sequence  int64
}

func NewSnowflake(workerId int64) (*Snowflake, error) {
	if workerId < 0 || workerId > workerIdMax {
		return nil, errors.New("workerId must be between 0 and 1023")
	}

	return &Snowflake{
		timestamp: 0,
		workerId:  workerId,
		sequence:  0,
	}, nil
}

func (sf *Snowflake) NextId() int64 {
	sf.Lock()
	defer sf.Unlock()

	now := time.Now().UnixNano() / 1e6

	if sf.timestamp == now {
		sf.sequence = (sf.sequence + 1) & sequenceMask

		if sf.sequence == 0 {
			for now <= sf.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		sf.sequence = 0
	}

	sf.timestamp = now

	nextId := (now-twepoch)<<timestampShift | (sf.workerId << workerIdShift) | (sf.sequence)

	return nextId
}
