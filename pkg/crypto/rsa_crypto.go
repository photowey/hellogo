package crypto

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
)

/**
 * crypto
 */

var (
	ErrPrivateKeyNULL = errors.New("the *rsa.PublicKey can't be nil")
	ErrPublicKeyNULL  = errors.New("the *rsa.PrivateKey can't be nil")
)

type CodecMode uint

var (
	PKCS1 CodecMode = 1 >> 1
	OAEP  CodecMode = 1 >> 2
)

// ------------------------------------------------------------------ Encrypt

// EncryptByPublicKeyOAEP
// RSA 公钥加密 OAEP
func EncryptByPublicKeyOAEP(data []byte, publicKey *rsa.PublicKey) (encrypted string, err error) {
	if publicKey == nil {
		return "", ErrPrivateKeyNULL
	}
	encryptedBytes, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, publicKey, data, nil)
	if err != nil {
		return "", fmt.Errorf("rsa:encrypt data err: %s", err.Error())
	}
	encrypted = encryptBase64(encryptedBytes)

	return encrypted, nil
}

// EncryptByPublicKeyPKCS1
// RSA 公钥加密 PKCS
func EncryptByPublicKeyPKCS1(data []byte, publicKey *rsa.PublicKey) (encrypted string, err error) {
	if publicKey == nil {
		return "", ErrPrivateKeyNULL
	}
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
	if err != nil {
		return "", fmt.Errorf("rsa:encrypt data err: %s", err.Error())
	}
	encrypted = encryptBase64(encryptedBytes)

	return encrypted, nil
}

// EncryptByPublicKeyBlockOAEP
// RSA 公钥块加密 OAEP
func EncryptByPublicKeyBlockOAEP(data []byte, publicKey *rsa.PublicKey) (encrypted string, err error) {
	if publicKey == nil {
		return "", ErrPrivateKeyNULL
	}

	encryptedBytes, err := blockEncrypt(data, publicKey, OAEP)
	if err != nil {
		return "", err
	}

	encrypted = encryptBase64(encryptedBytes)

	return encrypted, nil
}

// EncryptByPublicKeyBlockPKCS1
// RSA 公钥块加密 PKCS1
func EncryptByPublicKeyBlockPKCS1(data []byte, publicKey *rsa.PublicKey) (encrypted string, err error) {
	if publicKey == nil {
		return "", ErrPrivateKeyNULL
	}

	encryptedBytes, err := blockEncrypt(data, publicKey, PKCS1)
	if err != nil {
		return "", err
	}

	encrypted = encryptBase64(encryptedBytes)

	return encrypted, nil
}

// ------------------------------------------------------------------ Decrypt

// DecryptByPrivateKeyOAEP
// RSA 私钥解密 OAEP
func DecryptByPrivateKeyOAEP(encryptedBase64 string, privateKey *rsa.PrivateKey) (encrypted string, err error) {
	if privateKey == nil {
		return "", ErrPublicKeyNULL
	}

	encryptedBytes, err := decodeBase64(encryptedBase64)
	if err != nil {
		return "", err
	}

	encryptedByte, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, privateKey, encryptedBytes, nil)
	if err != nil {
		return "", fmt.Errorf("rsa:decrypt encrypted err:%s", err)
	}

	return string(encryptedByte), nil
}

// DecryptByPrivateKeyPKCS1
// RSA 私钥解密 PKCS1
func DecryptByPrivateKeyPKCS1(encryptedBase64 string, privateKey *rsa.PrivateKey) (encrypted string, err error) {
	if privateKey == nil {
		return "", ErrPublicKeyNULL
	}

	encryptedBytes, err := decodeBase64(encryptedBase64)
	if err != nil {
		return "", err
	}

	encryptedByte, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedBytes)
	if err != nil {
		return "", fmt.Errorf("rsa:decrypt encrypted err:%s", err)
	}

	return string(encryptedByte), nil
}

// DecryptByPrivateKeyBlockOAEP
// RSA 私钥块解密 OAEP
func DecryptByPrivateKeyBlockOAEP(encryptedBase64 string, privateKey *rsa.PrivateKey) (encrypted string, err error) {
	if privateKey == nil {
		return "", ErrPublicKeyNULL
	}

	encryptedBytes, err := decodeBase64(encryptedBase64)
	if err != nil {
		return "", err
	}
	// 分块处理
	decryptedBytes, err := blockDecrypt(privateKey, encryptedBytes, OAEP)
	if err != nil {
		return "", err
	}

	return string(decryptedBytes.Bytes()), err
}

// DecryptByPrivateKeyBlockPKCS1
// RSA 私钥块解密 PKCS1
func DecryptByPrivateKeyBlockPKCS1(encryptedBase64 string, privateKey *rsa.PrivateKey) (encrypted string, err error) {
	if privateKey == nil {
		return "", ErrPublicKeyNULL
	}

	encryptedBytes, err := decodeBase64(encryptedBase64)
	if err != nil {
		return "", err
	}
	// 分块处理
	decryptedBytes, err := blockDecrypt(privateKey, encryptedBytes, PKCS1)
	if err != nil {
		return "", err
	}

	return string(decryptedBytes.Bytes()), err
}

// ------------------------------------------------------------------ Block Codec

func blockEncrypt(data []byte, publicKey *rsa.PublicKey, codecMode CodecMode) ([]byte, error) {
	blockSize := publicKey.N.BitLen()/8 - 11
	blocks := split(data, blockSize)
	encryptedBytes := bytes.NewBufferString("")
	for _, block := range blocks {
		var encryptedByte []byte
		var err error
		if OAEP == codecMode {
			encryptedByte, err = rsa.EncryptOAEP(sha1.New(), rand.Reader, publicKey, block, nil)
		} else {
			encryptedByte, err = rsa.EncryptPKCS1v15(rand.Reader, publicKey, block)
		}
		if err != nil {
			return nil, err
		}

		encryptedBytes.Write(encryptedByte)
	}

	return encryptedBytes.Bytes(), nil
}

func blockDecrypt(privateKey *rsa.PrivateKey, encryptedBytes []byte, codecMode CodecMode) (*bytes.Buffer, error) {
	blockSize := privateKey.PublicKey.N.BitLen() / 8
	blocks := split(encryptedBytes, blockSize)
	decryptedBytes := bytes.NewBufferString("")
	for _, block := range blocks {

		var decryptedByte []byte
		var err error
		if OAEP == codecMode {
			decryptedByte, err = rsa.DecryptOAEP(sha1.New(), rand.Reader, privateKey, block, nil)
		} else {
			decryptedByte, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, block)
		}
		if err != nil {
			return nil, fmt.Errorf("rsa:decrypt data err:%s", err)
		}

		decryptedBytes.Write(decryptedByte)
	}

	return decryptedBytes, nil
}

// ------------------------------------------------------------------ Base64 Codec

func encryptBase64(encryptedBytes []byte) string {
	encrypted := base64.StdEncoding.EncodeToString(encryptedBytes)

	return encrypted
}

func decodeBase64(encryptedBase64 string) ([]byte, error) {
	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return nil, fmt.Errorf("base64:decode encryptedBytes data failed, error: %s", err.Error())
	}

	return encryptedBytes, nil
}
