package cron

import (
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/hellogo/pkg/collection"
	"go.uber.org/atomic"
)

type Interface interface {
	Run() error
	Stop()
	Add(*Task)
	Remove(string)
	Len() int
}

type Job interface {
	Name() string
	Exec() error
}

func SimpleJob(name string, exec func() error) Job {
	return &job{
		name: name,
		exec: exec,
	}
}

type job struct {
	name string
	exec func() error
}

func (j *job) Name() string {
	return j.name
}

func (j *job) Exec() error {
	return j.exec()
}

type Scheduler interface {
	Schedule(time.Time) time.Time
}

const RunAlways = -1

type schedule struct {
	from     time.Time
	times    int
	interval time.Duration
}

func (s *schedule) nextTime(t time.Time) time.Time {
	if s.from.IsZero() {
		return time.Time{}
	}

	if s.from.After(t) {
		return s.from
	}

	if s.interval == 0 {
		return time.Time{}
	}

	return t.Add(s.interval)
}

func (s *schedule) Schedule(t time.Time) time.Time {
	if s.times == 0 {
		return time.Time{}
	}

	next := s.nextTime(t)
	if !next.IsZero() {
		if s.times > 0 {
			s.times--
		}
	}

	return next
}

func Every(duration time.Duration) Scheduler {
	return &schedule{
		from:     time.Now(),
		times:    RunAlways,
		interval: duration,
	}
}

func EveryAt(from time.Time, duration time.Duration) Scheduler {
	return &schedule{
		from:     time.Now(),
		times:    RunAlways,
		interval: duration,
	}
}

func Once(at time.Time) Scheduler {
	return &schedule{
		from:  at,
		times: 1,
	}
}

func At(from time.Time, duration time.Duration, times int) Scheduler {
	return &schedule{
		from:     from,
		times:    times,
		interval: duration,
	}
}

type Task struct {
	Next time.Time
	Job
	Scheduler
}

type Cron struct {
	tasks   *collection.Heap[*Task]
	set     collection.Set[string]
	new     chan struct{}
	started *atomic.Bool
	logger  logr.Logger
}

func NewCron(logger logr.Logger) Interface {
	h := collection.NewHeap([]*Task{}, func(x, y *Task) bool {
		return x.Next.Before(y.Next)
	})

	return &Cron{
		tasks:   h,
		set:     collection.NewSet[string](),
		new:     make(chan struct{}, 8),
		started: new(atomic.Bool),
		logger:  logger,
	}
}

func (c *Cron) Add(task *Task) {
	if task.Next.IsZero() {
		task.Next = task.Scheduler.Schedule(time.Now())
	}
	c.tasks.Push(task)
	c.set.Insert(task.Name())

	if c.started.Load() {
		c.new <- struct{}{}
	}
}

func (c *Cron) Remove(name string) {
	c.set.Delete(name)
}

func (c *Cron) Len() int {
	return c.tasks.Len()
}

func (c *Cron) Run() error {
	c.started.Store(true)

	for {
		if !c.started.Load() {
			break
		}

		c.runTask()
	}

	return nil
}

const infTime time.Duration = 1<<63 - 1

func (c *Cron) runTask() {
	now := time.Now()
	duration := infTime
	task, ok := c.tasks.Peek()
	if ok {
		if !c.set.Has(task.Name()) {
			c.tasks.Pop()
			return
		}

		if task.Next.After(now) {
			duration = task.Next.Sub(now)
		} else {
			duration = 0
		}
	}

	timer := time.NewTimer(duration)
	defer timer.Stop()

	select {
	case <-c.new:
		return
	case <-timer.C:
	}

	task, ok = c.tasks.Pop()
	if !ok {
		return
	}

	go func() {
		start := time.Now()
		if err := task.Exec(); err != nil {
			c.logger.Info(fmt.Sprintf("Run job [%s] failed: %v", task.Name(), err))
			return
		}
		c.logger.Info(fmt.Sprintf("Run job [%s] successfully, duration %v", task.Name(), time.Since(start)))
	}()

	task.Next = task.Schedule(time.Now())
	if task.Next.IsZero() {
		c.set.Delete(task.Name())
	} else {
		c.tasks.Push(task)
	}
}

func (c *Cron) Stop() {
	c.started.Store(false)
	close(c.new)
}
