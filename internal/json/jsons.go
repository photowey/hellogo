package json

/**
 * jsons
 */

import (
	"encoding/json"
	"io"
)

// ---------------------------------------------------------------- JSON

func ToJSONString(body any) (string, error) {
	bytes, err := json.Marshal(body)

	return string(bytes), err
}

func ToJSONPretty(body any) (string, error) {
	bytes, err := json.MarshalIndent(body, "", "\t")

	return string(bytes), err
}

func ToStructd(reader io.Reader, target any) error {
	if err := json.NewDecoder(reader).Decode(target); err != nil {
		return err
	}

	return nil
}

func ToStruct(data []byte, target any) error {
	if err := json.Unmarshal(data, target); err != nil {
		return err
	}

	return nil
}
