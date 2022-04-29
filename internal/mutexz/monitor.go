package mutexz

//
// 监视器
//

import (
	"log"
	"sync"
	"time"

	"github.com/hellogo/internal/common/helper"
)

const (
	SelectTimeoutSeconds int64 = 3
)

// 初始化: 锁
var lock = &sync.Mutex{}

// 注册 handler: 锁
var handlerLock = &sync.Mutex{}

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

// Context 监视器处理上下文
type Context struct {
	Id      string         // 单次处理队列任务的标识
	Handler string         // 指定处理的 {@code handler} 名称
	Attrs   map[string]any // 参数列表
}

func NewContext() *Context {
	return &Context{
		Attrs: make(map[string]any),
	}
}

// Monitor 监视器
type Monitor struct {
	taskQueue   chan *Context
	resultQueue chan *Context
	registry    map[string]Handler
}

func NewMonitor() *Monitor {
	return &Monitor{
		taskQueue:   make(chan *Context),
		resultQueue: make(chan *Context),
		registry:    make(map[string]Handler),
	}
}

func (m *Monitor) Run() {
	for {
		select {
		case ctx := <-m.taskQueue:
			for _, handler := range m.registry {
				if handler.Supports(ctx.Handler) {
					handler.Handle(ctx)
					m.resultQueue <- ctx // 将 ctx 原路返回 各个 阻塞的 actor 根据自己的 Id 决策执行逻辑
				}
			}
			log.Printf("receive the monitor task,now:[%d]", time.Now().Second())
		case <-time.After(time.Duration(SelectTimeoutSeconds) * time.Second): // 3 s

		}
	}
}

func (m *Monitor) Submit(ctx *Context) {
	m.taskQueue <- ctx
}

// Handler {@code channel} 每次
type Handler interface {
	Supports(name string) bool
	Handle(ctx *Context)
}

// Register 注册 handler
func Register(name string, handler Handler) {
	handlerLock.Lock()
	defer handlerLock.Unlock()
	if helper.IsNil(handler) {
		panic("monitor: Register handler is nil")
	}
	if _, dup := monitor.registry[name]; dup {
		panic("monitor: Register called twice for handler " + name)
	}

	monitor.registry[name] = handler
}

func GetInstance() *Monitor {
	return monitor
}
