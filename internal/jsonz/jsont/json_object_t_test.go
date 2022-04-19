package jsont

import (
	`reflect`
	`testing`
)

func TestObject_Put(t *testing.T) {
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
			name: "Test Generic Object Put()",
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
			ob := &Object[int64]{
				ctx: tt.fields.context,
			}
			ob.Put(tt.args.key, tt.args.value)
		})
	}
}

func TestObject_Get(t *testing.T) {
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
			name: "Test Generic Object Get()-true",
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
			name: "Test Generic Object Get()-false",
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
			ob := &Object[int64]{
				ctx: tt.fields.context,
			}
			got, got1 := ob.Get(tt.args.key, tt.args.standBy)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.ok {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.ok)
			}
		})
	}
}

var body = `{
  "id": "9787111558422",
  "name": "The Go Programming Language",
  "authors": [
    "Alan A.A.Donovan",
    "Brian W. Kergnighan"
  ],
  "press": "Pearson Education"
}`

func TestParseObject(t *testing.T) {
	type args struct {
		body string
	}
	type testT[T any] struct {
		name string
		args args
		want *Object[T]
		ok   bool
	}

	tests := []testT[any]{
		{
			name: "Test Generic ParseObject()",
			args: args{
				body: body,
			},
			want: &Object[any]{},
			ok:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseObject[any](tt.args.body)
			if err != nil {
				t.Errorf("ParseObject() error = %v, ok %v", err, tt.ok)
				return
			}
			t.Logf("ParseObject() got = %v", got)
		})
	}
}
