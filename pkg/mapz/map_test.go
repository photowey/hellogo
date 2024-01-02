package mapz

import (
	"testing"
)

func TestHelloMap(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "hello map",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HelloMap()
		})
	}
}
