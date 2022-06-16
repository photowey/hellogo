package twentyhex

import (
	"testing"
)

func TestMustAlphabet(t *testing.T) {
	type args struct {
		alphabet string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test MustAlphabet-1",
			args: args{
				alphabet: "AZ",
			},
			want: true,
		},
		{
			name: "Test MustAlphabet-2",
			args: args{
				alphabet: "az",
			},
			want: true,
		},
		{
			name: "Test MustAlphabet-3",
			args: args{
				alphabet: "AZ1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustAlphabet(tt.args.alphabet); got != tt.want {
				t.Errorf("MustAlphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustTwentyHex(t *testing.T) {
	type args struct {
		alphabet string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test MustTwentyHex-1",
			args: args{
				alphabet: "0y00",
			},
			want: true,
		},
		{
			name: "Test MustTwentyHex-2",
			args: args{
				alphabet: "0y10",
			},
			want: true,
		},
		{
			name: "Test MustTwentyHex-3",
			args: args{
				alphabet: "0y1P",
			},
			want: true,
		},
		{
			name: "Test MustTwentyHex-4",
			args: args{
				alphabet: "1P",
			},
			want: false,
		},
		{
			name: "Test MustTwentyHex-4",
			args: args{
				alphabet: "0y1Q",
			},
			want: false,
		},
		{
			name: "Test MustTwentyHex-5",
			args: args{
				alphabet: "0y", // 0y00
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustTwentyHex(tt.args.alphabet); got != tt.want {
				t.Errorf("MustTwentyHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
