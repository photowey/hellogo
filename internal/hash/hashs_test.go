package hash

import (
	"testing"
)

func TestMakeHash(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "Test string make hash 32bit",
			args: args{
				str: "sharkchili",
			},
			want: 256048774,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeHash(tt.args.str); got != tt.want {
				t.Errorf("MakeHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
