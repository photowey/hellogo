package adapter

import (
	"testing"
)

func TestClient_InsertLightningConnectorIntoComputer(t *testing.T) {
	type args struct {
		com computer
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test insert light ning connector into computer",
			args: args{
				com: NewMac(),
			},
		},
		{
			name: "Test insert light ning connector into computer",
			args: args{
				com: NewWinAdapter(NewWin()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{}
			c.InsertLightningConnectorIntoComputer(tt.args.com)
		})
	}
}
