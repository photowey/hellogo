package condz

import (
	"testing"
)

func Test_run(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test sync.cond",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			run()
		})
	}
}
