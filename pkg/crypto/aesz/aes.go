package aesz

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/pkg/errors"
)

const (
	EmptyString = ""
)

type Context struct {
	Key string
	IV  string
}

func pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padTxt := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(src, padTxt...)
}

func unpad(src []byte) ([]byte, error) {
	length := len(src)
	if length == 0 {
		return []byte{0}, nil
	}
	unpadding := int(src[length-1])
	if unpadding > length {
		return nil, errors.New("unpad error.")
	}

	return src[:(length - unpadding)], nil
}

func Encrypt(data string, ctx Context) (string, error) {
	if EmptyString == data {
		return EmptyString, nil
	}
	block, err := aes.NewCipher([]byte(ctx.Key))
	if err != nil {
		return EmptyString, err
	}
	content := pad([]byte(data))

	cipherTxt := make([]byte, len(content))
	mode := cipher.NewCBCEncrypter(block, []byte(ctx.IV))
	mode.CryptBlocks(cipherTxt, content)
	encrypted := base64.StdEncoding.EncodeToString(cipherTxt)

	return encrypted, nil
}

func Decrypt(encryptedBase64 string, ctx Context) (string, error) {
	if EmptyString == encryptedBase64 {
		return EmptyString, nil
	}
	block, err := aes.NewCipher([]byte(ctx.Key))
	if err != nil {
		return EmptyString, err
	}
	decryptedBase64, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return EmptyString, err
	}
	if (len(decryptedBase64) % aes.BlockSize) != 0 {
		return "", errors.Errorf("block size:[%d] must be multiple of message length:[%d]", len(decryptedBase64), aes.BlockSize)
	}

	content := decryptedBase64
	mode := cipher.NewCBCDecrypter(block, []byte(ctx.IV))
	mode.CryptBlocks(content, content)
	decryptedBytes, err := unpad(content)
	if err != nil {
		return EmptyString, err
	}

	return string(decryptedBytes), nil
}
