package maps

import (
	"reflect"
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

func TestHashMap_Get(t *testing.T) {
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
		want   int64
		ok     bool
	}
	tests := []Test{
		{
			name: "Test HashMap#Get()-true",
			fields: fields[string, int64]{
				ctx: map[string]int64{
					"1": 1,
					"2": 2,
				},
			},
			args: args[string]{
				"2",
			},
			want: 2,
			ok:   true,
		},
		{
			name: "Test HashMap#Get()-false",
			fields: fields[string, int64]{
				ctx: map[string]int64{
					"1": 1,
					"2": 2,
				},
			},
			args: args[string]{
				"0",
			},
			want: 0,
			ok:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := &HashMap[string, int64]{
				ctx: tt.fields.ctx,
			}
			got, got1 := hm.Get(tt.args.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.ok {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.ok)
			}
		})
	}
}

func TestHashMap_Put(t *testing.T) {
	type fields[K comparable, V any] struct {
		ctx map[K]V
	}
	type args[K comparable, V any] struct {
		k K
		v V
	}
	type Test struct {
		name   string
		fields fields[string, int64]
		args   args[string, int64]
	}

	tests := []Test{
		{
			name: "Test HashMap#Put()",
			fields: fields[string, int64]{
				ctx: map[string]int64{
					"1": 1,
					"2": 2,
				},
			},
			args: args[string, int64]{
				k: "3",
				v: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := &HashMap[string, int64]{
				ctx: tt.fields.ctx,
			}
			hm.Put(tt.args.k, tt.args.v)
		})
	}
}

func TestHashMap_Remove(t *testing.T) {
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
		want   int64
	}

	tests := []Test{
		{
			name: "Test HashMap#Remove()-1",
			fields: fields[string, int64]{
				ctx: map[string]int64{
					"1": 1,
					"2": 2,
				},
			},
			args: args[string]{
				"1",
			},
			want: 1,
		},
		{
			name: "Test HashMap#Remove()-0",
			fields: fields[string, int64]{
				ctx: map[string]int64{
					"1": 1,
					"2": 2,
				},
			},
			args: args[string]{
				"3",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := &HashMap[string, int64]{
				ctx: tt.fields.ctx,
			}
			if got := hm.Remove(tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_PutAll(t *testing.T) {
	type fields[K comparable, V any] struct {
		ctx map[K]V
	}
	type args[K comparable, V any] struct {
		other Map[K, V]
	}

	type Test struct {
		name   string
		fields fields[string, int64]
		args   args[string, int64]
	}
	other := &HashMap[string, int64]{
		ctx: map[string]int64{
			"1": 1,
			"2": 2,
		},
	}
	tests := []Test{
		{
			name: "Test HashMap#PutAll()",
			fields: fields[string, int64]{
				ctx: make(map[string]int64, 0),
			},
			args: args[string, int64]{
				other: other,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := &HashMap[string, int64]{
				ctx: tt.fields.ctx,
			}
			hm.PutAll(tt.args.other)

			ex1, ok1 := hm.Get("1")
			t.Logf("the ex1:%d ok1:%T", ex1, ok1)

			ex2, ok2 := hm.Get("2")
			t.Logf("the ex2:%d ok2:%T", ex2, ok2)
		})
	}
}

func TestHashMap_Clear(t *testing.T) {
	type fields[K comparable, V any] struct {
		ctx map[K]V
	}

	type Test struct {
		name   string
		fields fields[string, int64]
	}

	tests := []Test{
		{
			name: "Test HashMap#Clear()",
			fields: fields[string, int64]{
				ctx: map[string]int64{
					"1": 1,
					"2": 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := &HashMap[string, int64]{
				ctx: tt.fields.ctx,
			}
			hm.Clear()
		})
	}
}

func TestHashMap_KeySet(t *testing.T) {
	type fields[K comparable, V any] struct {
		ctx map[K]V
	}

	type Test struct {
		name   string
		fields fields[string, int64]
		want   []string
	}

	tests := []Test{
		{
			name: "Test HashMap#KeySet()",
			fields: fields[string, int64]{
				ctx: map[string]int64{
					"1": 2,
				},
			},
			want: []string{"1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := &HashMap[string, int64]{
				ctx: tt.fields.ctx,
			}
			if got := hm.KeySet(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeySet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashMap_Values(t *testing.T) {
	type fields[K comparable, V any] struct {
		ctx map[K]V
	}

	type Test struct {
		name   string
		fields fields[string, int64]
		want   []int64
	}

	tests := []Test{
		{
			name: "Test HashMap#Values()",
			fields: fields[string, int64]{
				ctx: map[string]int64{
					"1": 2,
				},
			},
			want: []int64{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hm := &HashMap[string, int64]{
				ctx: tt.fields.ctx,
			}
			if got := hm.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}
