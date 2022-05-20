package collection

import (
	"reflect"
	"testing"
)

func TestInt64Keys(t *testing.T) {
	type args struct {
		ctx map[int64]int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "Test map<int64,int64> values",
			args: args{
				ctx: map[int64]int64{
					1024: 9527,
					6379: 8848,
					8761: 7923,
				},
			},
			want: []int64{1024, 6379, 8761},
		},
		{
			name: "Test map<int64,int64> values",
			args: args{
				ctx: map[int64]int64{},
			},
			want: []int64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64Keys(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}
