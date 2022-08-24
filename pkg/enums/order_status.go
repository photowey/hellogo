package enums

// $ go get golang.org/x/tools/cmd/stringer

//go:generate stringer -type=OrderStatus
type OrderStatus int

const (
	CREATE OrderStatus = iota + 1
	PAID
	DELIVERING
	COMPLETED
	CANCELLED
)

func OrderStatusEnumString(orderStatus OrderStatus) string {
	return orderStatus.String()
}

func OrderStatusEnumInt(orderStatus OrderStatus) int {
	return int(orderStatus)
}
