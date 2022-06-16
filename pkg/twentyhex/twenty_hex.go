package twentyhex

import (
	"math"
	"regexp"
	"strings"
)

var _ TwentyHex = (*twentyHex)(nil)

type TwentyHex interface {
	ToTwentyHex(decimal int) string
	ToInt(twentyHex string) int
	ToNext(twentyHex string) string
	ToNextAlphabet(twentyHex string) string
	ToIndex(ttyHex string) int
	ToAlphabet(twentyHex string) string
	ToAlphabetIndex(alphabet string) int
	FromAlphabet(alphabet string) string
}

type twentyHex struct {
}

func (ttyHex twentyHex) ToTwentyHex(decimal int) string {
	if decimal < 0 {
		return EmptyString
	}
	if decimal == 0 {
		return First
	}

	twentyHexx := EmptyString
	for decimal != 0 {
		index := decimal % Value
		ttHex := SystemAlphabet[index]
		twentyHexx = ttHex + twentyHexx
		decimal = decimal / Value
	}

	tail := twentyHexx
	if len(tail) == 1 {
		// 补位 - 高位补 0
		tail = SymbolHead + tail
	}

	return Symbol + tail
}

func (ttyHex twentyHex) ToInt(twentyHex string) int {
	decimal := 0
	if strings.TrimSpace(twentyHex) == EmptyString || twentyHex == Symbol {
		return decimal
	}
	if !MustTwentyHexAlphabet(twentyHex) {
		return -1
	}
	twentyHex = strings.ReplaceAll(twentyHex, Symbol, EmptyString)
	length := len(twentyHex)
	for i, next := range twentyHex {
		pow := float64((length - 1) - i)
		index := ttyHex.ToIndex(string(next))
		decimal += index * int(math.Pow(Value, pow))
	}

	return decimal
}

func (ttyHex twentyHex) ToNext(twentyHex string) string {
	if strings.TrimSpace(twentyHex) == EmptyString || !MustTwentyHex(twentyHex) {
		return EmptyString
	}
	idx := ttyHex.ToInt(twentyHex)

	return ttyHex.ToTwentyHex(idx + 1)
}

func (ttyHex twentyHex) ToNextAlphabet(twentyHex string) string {
	if strings.TrimSpace(twentyHex) == EmptyString || !MustTwentyHex(twentyHex) {
		return EmptyString
	}
	next := ttyHex.ToNext(twentyHex)

	return ttyHex.ToAlphabet(next)
}

func (ttyHex twentyHex) ToIndex(twentyHex string) int {
	for i := 0; i < len(SystemAlphabet); i++ {
		ttHex := SystemAlphabet[i]
		if ttHex == twentyHex {
			return i
		}
	}

	return -1
}

func (ttyHex twentyHex) ToAlphabet(twentyHex string) string {
	if strings.TrimSpace(twentyHex) == EmptyString || !MustTwentyHex(twentyHex) {
		return EmptyString
	}
	twentyHex = strings.ReplaceAll(twentyHex, Symbol, EmptyString)
	regex, _ := regexp.Compile("^0?")
	twentyHex = regex.ReplaceAllString(twentyHex, EmptyString)
	length := len(twentyHex)

	alphabet := EmptyString

	for i, next := range twentyHex {
		index := ttyHex.ToInt(string(next))
		if i != length-1 && index > 0 {
			index--
		}
		ttHex := Alphabet[index]

		alphabet += ttHex
	}

	return alphabet
}

func (ttyHex twentyHex) ToAlphabetIndex(alphabet string) int {
	if strings.TrimSpace(alphabet) == EmptyString || !MustAlphabet(alphabet) {
		return -1
	}

	for i := 0; i < len(Alphabet); i++ {
		alpha := Alphabet[i]
		if alpha == alphabet {
			return i
		}
	}

	return -1
}

func (ttyHex twentyHex) FromAlphabet(alphabet string) string {
	if strings.TrimSpace(alphabet) == EmptyString || !MustAlphabet(alphabet) {
		return EmptyString
	}
	length := len(alphabet)

	twentyHexx := EmptyString
	for i, next := range alphabet {
		index := ttyHex.ToAlphabetIndex(string(next))
		if i != (length-1) && index < len(SystemAlphabet)-1 {
			index++
		}
		ttHex := SystemAlphabet[index]
		twentyHexx += ttHex
	}

	tail := twentyHexx
	if len(tail) == 1 {
		tail = SymbolHead + tail
	}

	return Symbol + tail
}
