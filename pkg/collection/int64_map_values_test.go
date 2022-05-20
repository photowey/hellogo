package collection

import (
	"reflect"
	"testing"
)

func TestInt64Values(t *testing.T) {
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
			want: []int64{9527, 8848, 7923},
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
			if got := Int64Values(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64Values() = %v, want %v", got, tt.want)
			}
		})
	}
}
