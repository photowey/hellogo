package routine

import (
	"fmt"
	"testing"
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
			name: "Test goroutine worker",
			fields: fields{
				Fx: func() {
					fmt.Println("run in goroutine...")
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
}
