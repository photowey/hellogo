package jsont

import (
	`reflect`
	`testing`
)

func TestJSOONObject_Put(t *testing.T) {
	type fields[T any] struct {
		context map[string]T
	}
	type args[V any] struct {
		key   string
		value V
	}

	type testT[A any] struct {
		name   string
		fields fields[A]
		args   args[A]
	}

	tests := []testT[int64]{
		{
			name: "Test Generic JSOONObject Put()",
			fields: fields[int64]{
				context: make(map[string]int64),
			},
			args: args[int64]{
				key:   "hello",
				value: 1234567890987,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsoon := &JSOONObject[int64]{
				context: tt.fields.context,
			}
			jsoon.Put(tt.args.key, tt.args.value)
		})
	}
}

func TestJSOONObject_Get(t *testing.T) {
	type fields[T any] struct {
		context map[string]T
	}
	type args[A any] struct {
		key     string
		standBy A
	}
	type testT[V any] struct {
		name   string
		fields fields[V]
		args   args[V]
		want   V
		ok     bool
	}

	tests := []testT[int64]{
		{
			name: "Test Generic JSOONObject Get()-true",
			fields: fields[int64]{
				context: map[string]int64{
					"hello": 1234567890987,
				},
			},
			args: args[int64]{
				key:     "hello",
				standBy: 1234567890987,
			},
			want: 1234567890987,
			ok:   true,
		},
		{
			name: "Test Generic JSOONObject Get()-false",
			fields: fields[int64]{
				context: map[string]int64{
					"hello": 1234567890987,
				},
			},
			args: args[int64]{
				key:     "standBy",
				standBy: 0,
			},
			want: 0,
			ok:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsoon := &JSOONObject[int64]{
				context: tt.fields.context,
			}
			got, got1 := jsoon.Get(tt.args.key, tt.args.standBy)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.ok {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.ok)
			}
		})
	}
}
