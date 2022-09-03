package workerpool

import (
	"context"
)

type Task struct {
	ctx     context.Context
	handler Runnable
}

func NewTask(ctx context.Context, handler Runnable) Task {
	return Task{
		ctx:     ctx,
		handler: handler,
	}
}

func (task Task) Exec() {
	task.handler(task.ctx)
}

type Worker struct {
	WorkerQueue chan *Worker
	TaskChannel chan Task
	Stop        chan struct{}
}

func (w *Worker) Start() {
	go func() {
		var task Task
		for {
			w.WorkerQueue <- w
			select {
			case task = <-w.TaskChannel:
				task.Exec()
			case <-w.Stop:
				w.Stop <- struct{}{}
				return
			}
		}
	}()
}

func newWorker(pool chan *Worker) *Worker {
	return &Worker{
		WorkerQueue: pool,
		TaskChannel: make(chan Task),
		Stop:        make(chan struct{}),
	}
}

type Runnable func(ctx context.Context)

type Pool struct {
	WorkerQueue chan *Worker
	TaskQueue   chan Task
	stop        chan struct{}
}

func NewPool(workers uint, buffer uint) *Pool {
	workerQueue := make(chan *Worker, workers)
	taskQueue := make(chan Task, buffer)

	pool := &Pool{
		WorkerQueue: workerQueue,
		TaskQueue:   taskQueue,
		stop:        make(chan struct{}),
	}
	pool.Start()
	return pool
}

func (pool *Pool) Start() {
	for i := 0; i < cap(pool.WorkerQueue); i++ {
		worker := newWorker(pool.WorkerQueue)
		worker.Start()
	}

	go pool.dispatch()
}

func (pool *Pool) dispatch() {
	for {
		select {
		case task := <-pool.TaskQueue:
			worker := <-pool.WorkerQueue
			worker.TaskChannel <- task
		case <-pool.stop:
			for i := 0; i < cap(pool.WorkerQueue); i++ {
				worker := <-pool.WorkerQueue

				worker.Stop <- struct{}{}
				<-worker.Stop
			}

			pool.stop <- struct{}{}
			return
		}
	}
}

func (pool *Pool) Execute(task Task) {
	pool.TaskQueue <- task
}

func (pool *Pool) Release() {
	pool.stop <- struct{}{}
	<-pool.stop
}
