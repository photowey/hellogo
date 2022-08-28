package fanout

import (
	"context"
	"testing"
)

func TestFanout_Submit(t *testing.T) {
	type args struct {
		ctx     context.Context
		handler func(ctx context.Context)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Submit",
			args: args{
				ctx: context.Background(),
				handler: func(ctx context.Context) {
					panic("error")
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewFanout("testSubmit", WithWorker(1), WithBufferSize(1024))
			if err := queue.Submit(tt.args.ctx, tt.args.handler); (err != nil) != tt.wantErr {
				t.Errorf("Submit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFanout_Close(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Test Close",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := NewFanout("testClose", WithWorker(1), WithBufferSize(1024))
			if err := queue.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
