package crypto

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type RsaKey string

var (
	PRIVATE RsaKey = "PRIVATE KEY"
	PUBLIC  RsaKey = "PUBLIC KEY"
)

var (
	ErrPrivateKeyDecode = errors.New("decode the private key to block failed")
	ErrPublicKeyDecode  = errors.New("decode the public key to block failed")

	ErrPublicKeyParse    = errors.New("parse the public key failed")
	ErrInvalidPublicKey  = errors.New("invalid public key")
	ErrInvalidPrivateKey = errors.New("invalid private key")

	ErrInvalidPublicKeyType  = errors.New("invalid public key type")
	ErrInvalidPrivateKeyType = errors.New("invalid private key type")
)

// --------------------------------------------------------------------------------------

func LoadPublicKeyPem(publicKeyBytes []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil {
		return nil, ErrPublicKeyDecode
	}
	if block.Type != string(PUBLIC) {
		return nil, ErrInvalidPublicKeyType
	}

	return LoadPublicKey(block.Bytes)
}

func LoadPublicKey(blockByte []byte) (*rsa.PublicKey, error) {
	pubKey, err := x509.ParsePKIXPublicKey(blockByte)
	if err != nil {
		return nil, ErrPublicKeyParse
	}
	publicKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, ErrInvalidPublicKey
	}

	return publicKey, nil
}

func LoadPrivateKeyPem(privateKeyBytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		return nil, ErrPrivateKeyDecode
	}
	if block.Type != string(PRIVATE) {
		return nil, ErrInvalidPrivateKeyType
	}

	return LoadPrivateKey(block.Bytes)
}

func LoadPrivateKey(blockByte []byte) (*rsa.PrivateKey, error) {
	// {@code PKCS1}
	privateKey, err := x509.ParsePKCS1PrivateKey(blockByte)
	if err != nil {
		// Java 系 {@code PKCS8}
		pkcs8PrivateKey, err := x509.ParsePKCS8PrivateKey(blockByte)
		if err != nil {
			return nil, err
		}
		privateKey, ok := pkcs8PrivateKey.(*rsa.PrivateKey)
		if !ok {
			return nil, ErrInvalidPrivateKey
		}
		return privateKey, nil
	}

	return privateKey, err
}

// --------------------------------------------------------------------------------------

func split(data []byte, blockSize int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(data)/blockSize+1)
	for len(data) >= blockSize {
		chunk, data = data[:blockSize], data[blockSize:]
		chunks = append(chunks, chunk)
	}
	if len(data) > 0 {
		chunks = append(chunks, data)
	}

	return chunks
}
