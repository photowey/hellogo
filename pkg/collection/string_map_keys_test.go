package collection

import (
	"reflect"
	"testing"
)

func TestStringKeys(t *testing.T) {
	type args struct {
		ctx map[string]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test map<string,string> keys",
			args: args{
				ctx: map[string]string{
					"hello": "world",
					"tom":   "cat",
					"lilei": "hanmeimei",
				},
			},
			want: []string{"hello", "tom", "lilei"},
		},
		{
			name: "Test map<string,string> keys",
			args: args{
				ctx: map[string]string{},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringKeys(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
