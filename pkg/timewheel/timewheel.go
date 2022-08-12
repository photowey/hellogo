package timewheel

import (
	"container/list"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/hellogo/pkg/collection"
	"github.com/hellogo/pkg/cron"
)

type TimeWheel struct {
	interval    time.Duration
	slots       int
	currentSlot int
	tasks       []*list.List
	set         collection.Set[string]

	ticker *time.Ticker

	logger logr.Logger
}

type Task struct {
	cron.Task

	initialized bool
	slot        int
	circle      int
}

func NewTimeWheel(interval time.Duration, slots int, logger logr.Logger) cron.Interface {
	return &TimeWheel{
		interval: interval,
		slots:    slots,
		tasks:    make([]*list.List, slots),
		set:      collection.NewSet[string](),
		logger:   logger,
	}
}

func (tw *TimeWheel) Run() error {
	tw.ticker = time.NewTicker(tw.interval)

	for {
		now, ok := <-tw.ticker.C
		if !ok {
			break
		}
		tw.Exec(now, tw.currentSlot)
		tw.currentSlot = (tw.currentSlot + 1) % tw.slots
	}

	return nil
}

func (tw *TimeWheel) Exec(now time.Time, slot int) {
	taskList := tw.tasks[slot]
	if taskList == nil {
		return
	}

	for item := taskList.Front(); item != nil; {
		task, ok := item.Value.(*Task)
		if !ok || task == nil {
			item = item.Next()
			continue
		}

		if task.circle > 0 {
			task.circle--
			item = item.Next()
			continue
		}

		// run task
		go func() {
			start := time.Now()
			if err := task.Exec(); err != nil {
				tw.logger.Info(fmt.Sprintf("Run job [%s] failed: %v", task.Name(), err))
				return
			}
			tw.logger.Info(fmt.Sprintf("Run job [%s] successfully, duration %v", task.Name(), time.Since(start)))
		}()

		// delete or update task
		next := item.Next()
		taskList.Remove(item)
		item = next

		task.Next = task.Schedule(now)
		if !task.Next.IsZero() {
			tw.add(now, task)
		} else {
			tw.Remove(task.Name())
		}
	}
}

func (tw *TimeWheel) Stop() {
	tw.ticker.Stop()
}

func (tw *TimeWheel) Len() int {
	return tw.set.Len()
}

func (tw *TimeWheel) Add(task *cron.Task) {
	tw.add(time.Now(), &Task{
		Task: *task,
	})
}

func (tw *TimeWheel) add(now time.Time, task *Task) {
	if !task.initialized {
		task.Next = task.Schedule(now)
		task.initialized = true
	}

	duration := task.Next.Sub(now)
	if duration <= 0 {
		task.slot = tw.currentSlot + 1
		task.circle = 0
	} else {
		multi := int(duration / tw.interval)
		task.slot = (tw.currentSlot + multi) % tw.slots
		task.circle = multi / tw.slots
	}

	if tw.tasks[task.slot] == nil {
		tw.tasks[task.slot] = list.New()
	}

	tw.tasks[task.slot].PushBack(task)
	tw.set.Insert(task.Name())
}

func (tw *TimeWheel) Remove(name string) {
	tw.set.Delete(name)
}
