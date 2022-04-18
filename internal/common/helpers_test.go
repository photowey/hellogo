package common

import "testing"

func TestListFiles(t *testing.T) {
	type args struct {
		dir   string
		level int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test list dir",
			args: args{
				dir:   "/Users/photowey/Documents/playground/gopath/src/github.com/hellogo/internal",
				level: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListFiles(tt.args.dir, tt.args.level)
		})
	}
}
