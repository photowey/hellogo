package twentyhex

import (
	"testing"
)

func Test_twentyHex_FromAlphabet(t *testing.T) {
	type args struct {
		alphabet string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test FromAlphabet-1",
			args: args{
				alphabet: "A",
			},
			want: "0y00",
		},
		{
			name: "Test FromAlphabet-2",
			args: args{
				alphabet: "Z",
			},
			want: "0y0P",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ttyHex := twentyHex{}
			if got := ttyHex.FromAlphabet(tt.args.alphabet); got != tt.want {
				t.Errorf("FromAlphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twentyHex_ToAlphabet(t *testing.T) {
	type args struct {
		twentyHex string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test ToAlphabet-1",
			args: args{
				twentyHex: "0y00",
			},
			want: "A",
		},
		{
			name: "Test ToAlphabet-2",
			args: args{
				twentyHex: "0y0P",
			},
			want: "Z",
		},
		{
			name: "Test ToAlphabet-3",
			args: args{
				twentyHex: "0y10",
			},
			want: "AA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ttyHex := twentyHex{}
			if got := ttyHex.ToAlphabet(tt.args.twentyHex); got != tt.want {
				t.Errorf("ToAlphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twentyHex_ToAlphabetIndex(t *testing.T) {
	type args struct {
		alphabet string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test ToAlphabetIndex-1",
			args: args{
				alphabet: "A",
			},
			want: 0,
		},
		{
			name: "Test ToAlphabetIndex-2",
			args: args{
				alphabet: "Z",
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ttyHex := twentyHex{}
			if got := ttyHex.ToAlphabetIndex(tt.args.alphabet); got != tt.want {
				t.Errorf("ToAlphabetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twentyHex_ToIndex(t *testing.T) {
	type args struct {
		twentyHex string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test ToIndex-1",
			args: args{
				twentyHex: "0",
			},
			want: 0,
		},
		{
			name: "Test ToIndex-2",
			args: args{
				twentyHex: "P",
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ttyHex := twentyHex{}
			if got := ttyHex.ToIndex(tt.args.twentyHex); got != tt.want {
				t.Errorf("ToIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twentyHex_ToInt(t *testing.T) {
	type args struct {
		twentyHex string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test ToInt-1",
			args: args{
				twentyHex: "0y00",
			},
			want: 0,
		},
		{
			name: "Test ToInt-2",
			args: args{
				twentyHex: "0y10",
			},
			want: 26,
		},
		{
			name: "Test ToInt-3",
			args: args{
				twentyHex: "0y1P",
			},
			want: 51,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ttyHex := twentyHex{}
			if got := ttyHex.ToInt(tt.args.twentyHex); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twentyHex_ToNext(t *testing.T) {
	type args struct {
		twentyHex string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test ToNext-1",
			args: args{
				twentyHex: "0y00",
			},
			want: "0y01",
		},
		{
			name: "Test ToNext-2",
			args: args{
				twentyHex: "0y10",
			},
			want: "0y11",
		},
		{
			name: "Test ToNext-3",
			args: args{
				twentyHex: "0y1P",
			},
			want: "0y20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ttyHex := twentyHex{}
			if got := ttyHex.ToNext(tt.args.twentyHex); got != tt.want {
				t.Errorf("ToNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twentyHex_ToNextAlphabet(t *testing.T) {
	type args struct {
		twentyHex string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test ToNextAlphabet-1",
			args: args{
				twentyHex: "0y00",
			},
			want: "B",
		},
		{
			name: "Test ToNextAlphabet-2",
			args: args{
				twentyHex: "0y10",
			},
			want: "AB",
		},
		{
			name: "Test ToNextAlphabet-3",
			args: args{
				twentyHex: "0y1P",
			},
			want: "BA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ttyHex := twentyHex{}
			if got := ttyHex.ToNextAlphabet(tt.args.twentyHex); got != tt.want {
				t.Errorf("ToNextAlphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twentyHex_ToTwentyHex(t *testing.T) {
	type args struct {
		decimal int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test ToTwentyHex-1",
			args: args{
				decimal: 0,
			},
			want: "0y00",
		},
		{
			name: "Test ToTwentyHex-2",
			args: args{
				decimal: 26,
			},
			want: "0y10",
		},
		{
			name: "Test ToTwentyHex-3",
			args: args{
				decimal: 51,
			},
			want: "0y1P",
		},
		{
			name: "Test ToTwentyHex-4",
			args: args{
				decimal: 15123077923,
			},
			want: "0y1MOLL315",
		},
		{
			name: "Test ToTwentyHex-5",
			args: args{
				decimal: 10860143,
			},
			want: "0yNJN7L",
		},
		{
			name: "Test ToTwentyHex-6",
			args: args{
				decimal: 143, // 0x8F 0217 10001111
			},
			want: "0y5D",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ttyHex := twentyHex{}
			if got := ttyHex.ToTwentyHex(tt.args.decimal); got != tt.want {
				t.Errorf("ToTwentyHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
