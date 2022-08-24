package enums

import (
	"testing"
)

func TestOrderStatusEnumInt(t *testing.T) {
	type args struct {
		orderStatus OrderStatus
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test enum int",
			args: args{
				orderStatus: COMPLETED,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrderStatusEnumInt(tt.args.orderStatus); got != tt.want {
				t.Errorf("OrderStatusEnumInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderStatusEnumString(t *testing.T) {
	type args struct {
		orderStatus OrderStatus
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test enum string",
			args: args{
				orderStatus: COMPLETED,
			},
			want: "COMPLETED",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrderStatusEnumString(tt.args.orderStatus); got != tt.want {
				t.Errorf("OrderStatusEnumString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderStatus_String(t *testing.T) {
	tests := []struct {
		name string
		i    OrderStatus
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
