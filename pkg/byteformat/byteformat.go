package byteformat

import (
	"fmt"
)

const (
	Byte = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
)

func Format(bytes uint64) string {
	switch {
	case bytes < KB:
		return fmt.Sprintf("%dB", bytes)
	case bytes < MB:
		return fmt.Sprintf("%.2fK", float64(bytes)/KB)
	case bytes < GB:
		return fmt.Sprintf("%.2fM", float64(bytes)/MB)
	case bytes < TB:
		return fmt.Sprintf("%.2fG", float64(bytes)/GB)
	case bytes < PB:
		return fmt.Sprintf("%.2fT", float64(bytes)/TB)
	default:
		return fmt.Sprintf("%.2fP", float64(bytes)/PB)
	}
}
