package collection

import (
	"reflect"
	"testing"
)

func TestInt64AnyKeys(t *testing.T) {
	type args struct {
		ctx map[int64]any
	}

	type student struct {
		name string
	}

	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "Test map<int64,any> keys",
			args: args{
				ctx: map[int64]any{
					1024: "world",
					6379: 1024,
					8761: student{
						name: "hanmeimei",
					},
				},
			},
			want: []int64{1024, 6379, 8761},
		},
		{
			name: "Test map<int64,any> keys",
			args: args{
				ctx: map[int64]any{},
			},
			want: []int64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64AnyKeys(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64AnyKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
