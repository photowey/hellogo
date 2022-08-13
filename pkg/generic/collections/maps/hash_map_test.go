package maps

import (
	"testing"
)

func TestHashMap_Size(t *testing.T) {
	type fields[K comparable, V comparable] struct {
		ctx map[K]V
	}

	type Test struct {
		name   string
		fields fields[string, int64]
		want   int
	}

	tests := []Test{
		{
			name: "Test HashMap#Sze()-0",
			fields: fields[string, int64]{
				ctx: make(map[string]int64),
			},
			want: 0,
		}, {
			name: "Test HashMap#Sze()-1",
			fields: fields[string, int64]{
				ctx: map[string]int64{
					"1": 1,
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := HashMap[string, int64]{
				ctx: tt.fields.ctx,
			}
			if got := hm.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
