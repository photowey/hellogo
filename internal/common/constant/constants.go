package commonconstant

type (
	RequestSymbol     = string
	RequestCodeSymbol = int8
)

const (
	DecimalSystem        = 10 // 十进制
	DecimalSystemBitSize = 64 // 十进制-位数

	BooleanTrueInt64  = 1 // 真
	BooleanFalseInt64 = 0 // 假

	ZeroInt64       = BooleanFalseInt64 // 零
	EmptyCollection = BooleanFalseInt64 // 空列表
)

const (
	OneSecond         int64 = 1
	OneMinutesSeconds int64 = 60 * OneSecond
	OneHourSeconds    int64 = 60 * OneMinutesSeconds
	OneDaySeconds     int64 = OneHourSeconds * 24
)

const (
	DefaultEmptyString string = "" // 空串
)
