package common

import (
	"sort"
	"strings"
)

var (
	DefaultSeparator string = ","
	SignSeparator    string = "&"
	DefaultJoiner    string = "="
)

type StringBuffer struct {
	buffer []string
}

func (sb StringBuffer) Append(needle string) StringBuffer {
	_ = append(sb.buffer, needle)

	return sb
}

func (sb StringBuffer) Join(key, value, joiner string) string {
	return key + joiner + value
}

func (sb StringBuffer) ToString() string {
	return implode(sb.buffer, SignSeparator)
}

func (sb StringBuffer) ToStrings(separator string) string {
	return implode(sb.buffer, separator)
}

func (sb StringBuffer) ToSortString() string {
	cloneSlice := sb.cloneSlice()
	sort.Strings(cloneSlice)

	return implode(cloneSlice, SignSeparator)
}

func (sb StringBuffer) ToSortStrings(separator string) string {
	cloneSlice := sb.cloneSlice()
	sort.Strings(cloneSlice)

	return implode(cloneSlice, separator)
}

func (sb StringBuffer) cloneSlice() []string {
	cloneSlice := make([]string, len(sb.buffer))
	copy(cloneSlice, sb.buffer)

	return cloneSlice
}

func implode(haystack []string, separator string) string {
	var buf strings.Builder
	for _, str := range haystack {
		buf.WriteString(str)
		buf.WriteString(separator)
	}

	return strings.TrimRight(buf.String(), separator)
}
