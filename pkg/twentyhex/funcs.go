package twentyhex

import (
	"regexp"
)

const (
	AlphabetRegex       = "^[A-Za-z]*$"
	SystemRegex         = "^0y[A-P0-9]*$"
	SystemAlphabetRegex = "[A-P0-9]*$"
)

func MustMatch(pattern, alphabet string) bool {
	ok, _ := regexp.MatchString(pattern, alphabet)

	return ok
}

func MustAlphabet(alphabet string) bool {
	return MustMatch(AlphabetRegex, alphabet)
}

func MustTwentyHex(alphabet string) bool {
	return MustMatch(SystemRegex, alphabet)
}

func MustTwentyHexAlphabet(alphabet string) bool {
	return MustMatch(SystemAlphabetRegex, alphabet)
}
