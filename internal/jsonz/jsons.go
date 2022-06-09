package jsonz

/**
 * jsonz
 */

import (
	"encoding/json"
	"io"
)

func String(body any) (string, error) {
	bytez, err := json.Marshal(body)

	return string(bytez), err
}

func Pretty(body any) (string, error) {
	bytez, err := json.MarshalIndent(body, "", "\t")

	return string(bytez), err
}

func DecodeStruct(reader io.Reader, structy any) error {
	if err := json.NewDecoder(reader).Decode(structy); err != nil {
		return err
	}

	return nil
}

func UnmarshalStruct(data []byte, structy any) error {
	if err := json.Unmarshal(data, structy); err != nil {
		return err
	}

	return nil
}

func UnmarshalMap(data []byte) (map[string]string, error) {
	maps := make(map[string]string)
	if err := json.Unmarshal(data, &maps); err != nil {
		return maps, err
	}

	return maps, nil
}