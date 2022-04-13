package common

import (
	"bytes"
	"io"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func ToUTF8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	dataBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return dataBytes, nil
}

func ToGBK(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	dataBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return dataBytes, nil
}
