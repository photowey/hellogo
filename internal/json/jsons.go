package json

/**
 * jsons
 */

import (
	"encoding/json"
	"io"
)

func ToJSONString(body any) (string, error) {
	bytez, err := json.Marshal(body)

	return string(bytez), err
}

func ToJSONPretty(body any) (string, error) {
	bytez, err := json.MarshalIndent(body, "", "\t")

	return string(bytez), err
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
