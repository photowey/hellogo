package async_test

import (
	"testing"
	"time"

	"github.com/hellogo/internal/async"
)

func TestFuture(t *testing.T) {
	t.Log("Let's start ...")
	future := async.Run(func() any {
		return DoneAsync(t)
	})
	t.Log("Done is running ...")
	futureValue := future.Await()
	t.Log(futureValue)
}

func DoneAsync(t *testing.T) int {
	t.Log("Warming up ...")
	time.Sleep(3 * time.Second)
	t.Log("Done ...")

	return 1
}
