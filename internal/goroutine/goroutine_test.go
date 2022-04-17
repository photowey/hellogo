package goroutine

import (
	"testing"
	"time"
)

func TestGoroutine_Run(t *testing.T) {
	type fields struct {
		Fx func()
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Test goroutine worker run",
			fields: fields{
				Fx: func() {
					t.Log("run in goroutine...")
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actor := Goroutine{
				Fx: tt.fields.Fx,
			}
			actor.Run()
		})
	}

	time.Sleep(time.Duration(1) * time.Second)
}

func TestGoroutine_Runp(t *testing.T) {
	type fields struct {
		Fxp func(params any)
	}

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Test goroutine worker run with parameters",
			fields: fields{
				Fxp: func(params any) {
					t.Logf("run in goroutine with parameters:[%s]...", params)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actor := Goroutine{
				Fxp: tt.fields.Fxp,
			}
			actor.Runo()
			actor.Runo("hello goroutine")
			actor.Runo("hello ", "world")
		})
	}

	time.Sleep(time.Duration(1) * time.Second)
}

func TestGoroutine_Runo(t *testing.T) {
	type fields struct {
		Fxo func(params ...any)
	}

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Test goroutine worker run with parameters",
			fields: fields{
				Fxo: func(params ...any) {
					t.Logf("run in goroutine with parameters: %v", params...)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actor := Goroutine{
				Fxo: tt.fields.Fxo,
			}
			actor.Runo()
			actor.Runo("hello,goroutine")
			actor.Runo("hello", "world")
		})
	}

	time.Sleep(time.Duration(1) * time.Second)
}
