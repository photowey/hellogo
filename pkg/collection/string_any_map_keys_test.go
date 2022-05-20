package collection

import (
	"reflect"
	"testing"
)

func TestStringAnyKeys(t *testing.T) {
	type args struct {
		ctx map[string]any
	}

	type student struct {
		name string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test map<string,any> keys",
			args: args{
				ctx: map[string]any{
					"hello": "world",
					"tom":   1024,
					"lilei": student{
						name: "hanmeimei",
					},
				},
			},
			want: []string{"hello", "tom", "lilei"},
		},
		{
			name: "Test map<string,any> keys",
			args: args{
				ctx: map[string]any{},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringAnyKeys(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringAnyKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
