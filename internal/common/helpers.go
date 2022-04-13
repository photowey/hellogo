package common

import (
	"os"
)

func Match(expect bool, standard, defaultCandidate string) string {
	if expect {
		return standard
	}
	return defaultCandidate
}

func Exists(path string) bool {
	_, err := os.Stat(path)

	return err != nil || os.IsExist(err)
}
