package castz

import (
	"reflect"
	"testing"
	"time"
)

func TestToBoolB(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToBoolB(tt.args.src)
			if got != tt.want {
				t.Errorf("ToBoolB() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToBoolB() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToFloat32B(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  float32
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToFloat32B(tt.args.src)
			if got != tt.want {
				t.Errorf("ToFloat32B() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToFloat32B() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToFloat64B(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToFloat64B(tt.args.src)
			if got != tt.want {
				t.Errorf("ToFloat64B() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToFloat64B() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToInt16B(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  int16
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToInt16B(tt.args.src)
			if got != tt.want {
				t.Errorf("ToInt16B() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToInt16B() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToInt32B(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  int32
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToInt32B(tt.args.src)
			if got != tt.want {
				t.Errorf("ToInt32B() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToInt32B() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToInt64B(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToInt64B(tt.args.src)
			if got != tt.want {
				t.Errorf("ToInt64B() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToInt64B() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToInt8B(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  int8
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToInt8B(tt.args.src)
			if got != tt.want {
				t.Errorf("ToInt8B() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToInt8B() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToIntB(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToIntB(tt.args.src)
			if got != tt.want {
				t.Errorf("ToIntB() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToIntB() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToIntSliceB(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToIntSliceB(tt.args.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToIntSliceB() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToIntSliceB() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToSliceB(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  []any
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToSliceB(tt.args.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSliceB() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToSliceB() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToStringB(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToStringB(tt.args.src)
			if got != tt.want {
				t.Errorf("ToStringB() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToStringB() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToStringMapB(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]any
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToStringMapB(tt.args.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringMapB() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToStringMapB() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToStringMapInt64B(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]int64
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToStringMapInt64B(tt.args.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringMapInt64B() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToStringMapInt64B() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToStringMapStringB(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]string
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToStringMapStringB(tt.args.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringMapStringB() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToStringMapStringB() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToStringMapStringSliceB(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  map[string][]string
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToStringMapStringSliceB(tt.args.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringMapStringSliceB() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToStringMapStringSliceB() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToStringSliceB(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToStringSliceB(tt.args.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringSliceB() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToStringSliceB() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToTimeB(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  time.Time
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToTimeB(tt.args.src)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToTimeB() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToTimeB() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToUInt16B(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  uint16
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToUInt16B(tt.args.src)
			if got != tt.want {
				t.Errorf("ToUInt16B() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToUInt16B() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToUInt32B(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  uint32
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToUInt32B(tt.args.src)
			if got != tt.want {
				t.Errorf("ToUInt32B() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToUInt32B() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToUInt64B(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  uint64
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToUInt64B(tt.args.src)
			if got != tt.want {
				t.Errorf("ToUInt64B() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToUInt64B() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToUInt8B(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  uint8
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToUInt8B(tt.args.src)
			if got != tt.want {
				t.Errorf("ToUInt8B() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToUInt8B() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToUIntB(t *testing.T) {
	type args struct {
		src any
	}
	tests := []struct {
		name  string
		args  args
		want  uint
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToUIntB(tt.args.src)
			if got != tt.want {
				t.Errorf("ToUIntB() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToUIntB() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
