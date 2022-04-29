package helper

import (
	"bytes"
	"runtime"
	"strconv"

	"github.com/gomodule/redigo/redis"

	commonconstant "github.com/hellogo/internal/common/constant"
)

// MakeIntSlice 初始化一个指定容量的 {@code int} {@code slice}
func MakeIntSlice(length int) []int64 {
	return make([]int64, length)
}

// InitIntSlice 初始化一个 {@code int64} {@code slice}
func InitIntSlice(opts ...int64) []int64 {
	slice := MakeIntSlice(len(opts))
	for i, opt := range opts {
		slice[i] = opt
	}

	return slice
}

// MakeIntMap 文档戎商 见{@code MakeStringMap}
func MakeIntMap(cap ...int) map[string]int64 {
	switch len(cap) {
	case 0:
		return make(map[string]int64)
	case 1:
		return make(map[string]int64, cap[0])
	default:
		return make(map[string]int64)
	}
}

// ArrayIntContains 判定一个整型切片中是否包含另外一个整型元素
func ArrayIntContains(haystack []int64, needle int64) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}

	return false
}

// StringToInt64 将数值类型的字符串转换为 {@code int64}
func StringToInt64(needle string) (int64, error) {
	return strconv.ParseInt(needle, commonconstant.DecimalSystem, commonconstant.DecimalSystemBitSize)
}

// StringSliceToInt64 将数值类型的字符串切片转换为 {@code int64} 类型切片
func StringSliceToInt64(haystack []string) ([]int64, error) {
	if len(haystack) == 0 {
		return MakeIntSlice(0), nil
	}
	haystacks := MakeIntSlice(len(haystack))
	for i, v := range haystack {
		// _, 进制, 位数
		value, err := StringToInt64(v)
		if err != nil {
			return MakeIntSlice(0), err
		}

		haystacks[i] = value
	}

	return haystacks, nil
}

// ---------------------------------------------------------------- int64

// FormatInt64 格式化 int64 为字符串
func FormatInt64(needle int64) string {
	return strconv.FormatInt(needle, commonconstant.DecimalSystem)
}

// FormatFloat64 将字符串格式化为 {@code float64}
func FormatFloat64(needle string) (float64, error) {
	fv, err := strconv.ParseFloat(needle, commonconstant.DecimalSystemBitSize)

	return fv, err
}

// ---------------------------------------------------------------- uint64

// GetGoroutineId 获取当前执行的 {@code Goroutine} 底层 AppId
func GetGoroutineId() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)

	return n
}

// ---------------------------------------------------------------- return bool

func IsBlankString(str string) bool {
	return "" == str
}

func IsNotBlankString(str string) bool {
	return !IsBlankString(str)
}

func IsNil(needle any) bool {
	return nil == needle
}

func IsNotNil(needle any) bool {
	return !IsNil(needle)
}

func IsErrorNil(err error) bool {
	return nil == err
}

func IsNotNilError(err error) bool {
	return !IsErrorNil(err)
}

// IsRedisNilError 当 {@code Redis} 中没有指定 key 的时候, 就会返回 {@code redis.ErrNil}
//
// 开发是需要特别注意
func IsRedisNilError(err error) bool {
	return err.Error() == redis.ErrNil.Error()
}

// IsEmptyCollection 空集合
//
// 1.数组
//
// 2.切片
//
// 3.map
//
// 4....
func IsEmptyCollection(target []any) bool {
	return len(target) == 0
}

// IsNotEmptyCollection 非空集合
func IsNotEmptyCollection(target []any) bool {
	return !IsEmptyCollection(target)
}
