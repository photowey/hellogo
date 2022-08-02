package hicooflake

import (
	"testing"
)

func TestHicooflake_NextId(t *testing.T) {
	tests := []struct {
		name string
		want uint64
	}{
		{
			name: "Test nextId",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hf := NowHicooflake()
			if got := hf.NextId(); got <= tt.want {
				t.Errorf("NextId() = %v, want %v", got, tt.want)
			}
		})
	}
}
