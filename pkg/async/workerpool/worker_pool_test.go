package workerpool

import (
	"context"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

func init() {
	nCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(nCPUs)
}

func TestNewPool(t *testing.T) {
	type args struct {
		workers uint
		buffer  uint
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test NewPool",
			args: args{
				workers: 1000,
				buffer:  10_000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pool := NewPool(tt.args.workers, tt.args.buffer)
			defer pool.Release()

			times := 1_000_000
			var counter uint64

			wg := sync.WaitGroup{}
			wg.Add(times)
			for i := 0; i < times; i++ {
				arg := uint64(1)
				task := NewTask(context.Background(), func(ctx context.Context) {
					defer wg.Done()
					atomic.AddUint64(&counter, arg)
				})

				pool.Execute(task)
			}
			wg.Wait()

			counterFinal := atomic.LoadUint64(&counter)
			if uint64(times) != counterFinal {
				t.Errorf("times %v is not equal counterFinal %v", times, counterFinal)
			}
		})
	}
}
