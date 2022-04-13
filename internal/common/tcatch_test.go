package common

import (
	"testing"
)

func TestTryCatcher_Try(t *testing.T) {
	message := "404"
	Try(func() {
		panic(message)
	}).Catch("not found", func(err any) {
		t.Log("handle not found exception")
	}).Catch("404", func(err any) {
		t.Log("handle 404 exception")
	}).Finally(func() {
		t.Log("handle finally")
	})
}
