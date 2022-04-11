package common

import (
	"encoding/json"
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

// ---------------------------------------------------------------- Contain

func StringsContains(array []string, val string) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

// ---------------------------------------------------------------- JSON

func ToJSONObject(body any) (string, error) {
	bytes, err := json.Marshal(body)

	return string(bytes), err
}

func ToJSONPretty(body any) (string, error) {
	bytes, err := json.MarshalIndent(body, "", "\t")

	return string(bytes), err
}
