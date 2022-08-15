package maps

import (
	"testing"
)

func TestHashMap_Size(t *testing.T) {
	type fields[K comparable, V any] struct {
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

func TestHashMap_IsEmpty(t *testing.T) {
	type fields[K comparable, V any] struct {
		ctx map[K]V
	}

	type Test struct {
		name   string
		fields fields[string, int64]
		want   bool
	}

	tests := []Test{
		{
			name: "Test HashMap#IsEmpty()-true",
			fields: fields[string, int64]{
				ctx: make(map[string]int64),
			},
			want: true,
		}, {
			name: "Test HashMap#IsEmpty()-false",
			fields: fields[string, int64]{
				ctx: map[string]int64{
					"1": 1,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := HashMap[string, int64]{
				ctx: tt.fields.ctx,
			}
			if got := hm.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_ContainsKey(t *testing.T) {
	type fields[K comparable, V any] struct {
		ctx map[K]V
	}

	type args[K comparable] struct {
		k K
	}

	type Test struct {
		name   string
		fields fields[string, int64]
		args   args[string]
		want   bool
	}

	tests := []Test{
		{
			name: "Test HashMap#ContainsKey()-false",
			fields: fields[string, int64]{
				ctx: make(map[string]int64, 0),
			},
			args: args[string]{
				"1",
			},
			want: false,
		},
		{
			name: "Test HashMap#ContainsKey()-true",
			fields: fields[string, int64]{
				ctx: map[string]int64{
					"1": 1,
					"2": 2,
				},
			},
			args: args[string]{
				"2",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := HashMap[string, int64]{
				ctx: tt.fields.ctx,
			}
			if got := hm.ContainsKey(tt.args.k); got != tt.want {
				t.Errorf("ContainsKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_ContainsValue(t *testing.T) {
	type fields[K comparable, V any] struct {
		ctx map[K]V
	}

	type args[V any] struct {
		v V
	}

	type Test struct {
		name   string
		fields fields[string, int64]
		args   args[int64]
		want   bool
	}
	tests := []Test{
		{
			name: "Test HashMap#ContainsValue()-false",
			fields: fields[string, int64]{
				ctx: make(map[string]int64, 0),
			},
			args: args[int64]{
				1,
			},
			want: false,
		},
		{
			name: "Test HashMap#ContainsValue()-false",
			fields: fields[string, int64]{
				ctx: map[string]int64{
					"1": 1,
					"2": 2,
				},
			},
			args: args[int64]{
				0,
			},
			want: false,
		},
		{
			name: "Test HashMap#ContainsValue()-true",
			fields: fields[string, int64]{
				ctx: map[string]int64{
					"1": 1,
					"2": 2,
				},
			},
			args: args[int64]{
				2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := HashMap[string, int64]{
				ctx: tt.fields.ctx,
			}
			if got := hm.ContainsValue(tt.args.v); got != tt.want {
				t.Errorf("ContainsValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
