package mutexz

import (
	`log`
	`sync`
	`time`
)

var lock = &sync.Mutex{}

var monitor *Monitor

func init() {
	if nil == monitor {
		lock.Lock()
		defer lock.Unlock()
		if nil == monitor {
			monitor = NewMonitor()
		}
	}
}

type Symbol struct {
}

func NewWatch() Symbol {
	return Symbol{}
}

type Monitor struct {
	taskQueue   chan Symbol
	resultQueue chan Symbol
}

func NewMonitor() *Monitor {
	return &Monitor{
		taskQueue:   make(chan Symbol),
		resultQueue: make(chan Symbol),
	}
}

func (m *Monitor) Run() {
	for {
		select {
		case <-m.taskQueue:
			log.Printf("receive the monitor task,now:[%d]", time.Now().Second())
		}
	}
}

func (m *Monitor) Submit() {
	m.taskQueue <- NewWatch()
}

func GetInstance() *Monitor {
	return monitor
}
