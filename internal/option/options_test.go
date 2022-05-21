package option

import (
	"reflect"
	"testing"
)

func TestOptionalEmpty(t *testing.T) {
	type args[T any] struct {
		zero T
	}

	type testT[A any] struct {
		name string
		args args[A]
		want Optional[A]
	}

	tests := []testT[string]{
		{
			name: "Test string optional empty()",
			args: args[string]{
				zero: "",
			},
			want: Optional[string]{
				data:      "",
				present:   false,
				valueType: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OptionalEmpty(tt.args.zero); got.IsPresent() || !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OptionalEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptionalOf(t *testing.T) {
	type args[T any] struct {
		value T
	}

	type testT[A any] struct {
		name string
		args args[A]
		want Optional[A]
	}

	tests := []testT[string]{
		{
			name: "Test string optional of()",
			args: args[string]{
				value: "hello",
			},
			want: Optional[string]{
				data:      "hello",
				present:   true,
				valueType: "string",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OptionalOf(tt.args.value); !got.IsPresent() || !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OptionalOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptional_String(t *testing.T) {
	type fields[T any] struct {
		Data      T
		Present   bool
		ValueType string
	}
	type testT[A any] struct {
		name   string
		fields fields[A]
		want   string
	}

	// 测试 string
	tests := []testT[string]{
		{
			name: "Test string optional string()",
			fields: fields[string]{
				Data:      "hello",
				Present:   true,
				ValueType: "string",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optional := Optional[string]{
				data:      tt.fields.Data,
				present:   tt.fields.Present,
				valueType: tt.fields.ValueType,
			}
			if got := optional.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}

	// 测试 int
	testInts := []testT[int64]{
		{
			name: "Test int optional string()",
			fields: fields[int64]{
				Data:      9527884879236379,
				Present:   true,
				ValueType: "int64",
			},
			want: "9527884879236379",
		},
	}
	for _, tt := range testInts {
		t.Run(tt.name, func(t *testing.T) {
			optional := Optional[int64]{
				data:      tt.fields.Data,
				present:   tt.fields.Present,
				valueType: tt.fields.ValueType,
			}
			if got := optional.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptional_Get(t *testing.T) {
	type fields[T any] struct {
		Data      T
		Present   bool
		ValueType string
	}
	type testT[A any] struct {
		name   string
		fields fields[A]
		want   A
	}

	// 测试 string
	tests := []testT[string]{
		{
			name: "Test string optional string()",
			fields: fields[string]{
				Data:      "hello",
				Present:   true,
				ValueType: "string",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optional := Optional[string]{
				data:      tt.fields.Data,
				present:   tt.fields.Present,
				valueType: tt.fields.ValueType,
			}
			if got := optional.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}

	// 测试 int
	testInts := []testT[int64]{
		{
			name: "Test int optional string()",
			fields: fields[int64]{
				Data:      9527884879236379,
				Present:   true,
				ValueType: "int64",
			},
			want: 9527884879236379,
		},
	}

	for _, tt := range testInts {
		t.Run(tt.name, func(t *testing.T) {
			optional := Optional[int64]{
				data:      tt.fields.Data,
				present:   tt.fields.Present,
				valueType: tt.fields.ValueType,
			}
			if got := optional.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptional_OrElse(t *testing.T) {
	type fields[T any] struct {
		Data      T
		Present   bool
		ValueType string
	}
	type args[A any] struct {
		standBy A
	}

	type testT[D any] struct {
		name   string
		fields fields[D]
		args   args[D]
		want   D
	}

	tests := []testT[string]{
		{
			name: "Test string optional OrElse()",
			fields: fields[string]{
				Data:      "",
				Present:   false,
				ValueType: "",
			},
			args: args[string]{
				standBy: "I'm standBy parameter",
			},
			want: "I'm standBy parameter",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optional := Optional[string]{
				data:      tt.fields.Data,
				present:   tt.fields.Present,
				valueType: tt.fields.ValueType,
			}
			if got := optional.OrElse(tt.args.standBy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrElse() = %v, want %v", got, tt.want)
			}
		})
	}

	testPresents := []testT[string]{
		{
			name: "Test string optional OrElse()",
			fields: fields[string]{
				Data:      "Present",
				Present:   true,
				ValueType: "string",
			},
			args: args[string]{
				standBy: "I'm standBy parameter",
			},
			want: "Present",
		},
	}
	for _, tt := range testPresents {
		t.Run(tt.name, func(t *testing.T) {
			optional := Optional[string]{
				data:      tt.fields.Data,
				present:   tt.fields.Present,
				valueType: tt.fields.ValueType,
			}
			if got := optional.OrElse(tt.args.standBy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrElse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComparableOptional_Equals(t *testing.T) {
	type fields[T any] struct {
		Optional Optional[T]
	}

	type args[A any] struct {
		value A
	}

	type testT[D any] struct {
		name   string
		fields fields[D]
		args   args[D]
		want   bool
	}

	tests := []testT[int64]{
		{
			name: "Test comparable option equals",
			fields: fields[int64]{
				Optional: Optional[int64]{
					data:      9527,
					present:   true,
					valueType: "int64",
				},
			},
			args: args[int64]{
				value: 9527,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optional := ComparableOptional[int64]{
				Optional: tt.fields.Optional,
			}
			if got := optional.Equals(tt.args.value); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
