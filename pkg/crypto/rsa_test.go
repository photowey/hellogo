package crypto

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"strings"
	"testing"
)

var rsaPrivateKey = &rsa.PrivateKey{
	PublicKey: rsa.PublicKey{
		N: fromBase10("9353930466774385905609975137998169297361893554149986716853295022578535724979677252958524466350471210367835187480748268864277464700638583474144061408845077"),
		E: 65537,
	},
	D: fromBase10("7266398431328116344057699379749222532279343923819063639497049039389899328538543087657733766554155839834519529439851673014800261285757759040931985506583861"),
	Primes: []*big.Int{
		fromBase10("98920366548084643601728869055592650835572950932266967461790948584315647051443"),
		fromBase10("94560208308847015747498523884063394671606671904944666360068158221458669711639"),
	},
}

type signPKCS1v15Test struct {
	in, out string
}

// These vectors have been tested with
//   `openssl rsautl -verify -inkey pk -in signature | hexdump -C`
var signPKCS1v15Tests = []signPKCS1v15Test{
	{"Test.\n", "a4f3fa6ea93bcdd0c57be020c1193ecbfd6f200a3d95c409769b029578fa0e336ad9a347600e40d3ae823b8c7e6bad88cc07c1d54c3a1523cbbb6d58efc362ae"},
}

func fromBase10(base10 string) *big.Int {
	i, ok := new(big.Int).SetString(base10, 10)
	if !ok {
		panic("bad number: " + base10)
	}
	return i
}

func TestSignPKCS1v15(t *testing.T) {
	for i, test := range signPKCS1v15Tests {
		h := sha256.New()
		h.Write([]byte(test.in))
		digest := h.Sum(nil)

		s, err := rsa.SignPKCS1v15(nil, rsaPrivateKey, crypto.SHA256, digest)
		if err != nil {
			t.Errorf("#%d %s", i, err)
		}

		expected, _ := hex.DecodeString(test.out)
		if !bytes.Equal(s, expected) {
			t.Errorf("#%d got: %x want: %x", i, s, expected)
		}
	}
}

func TestVerifyPKCS1v15(t *testing.T) {
	for i, test := range signPKCS1v15Tests {
		h := sha256.New()
		h.Write([]byte(test.in))
		digest := h.Sum(nil)

		sig, _ := hex.DecodeString(test.out)

		err := rsa.VerifyPKCS1v15(&rsaPrivateKey.PublicKey, crypto.SHA256, digest, sig)
		if err != nil {
			t.Errorf("#%d %s", i, err)
		}
	}
}

func TestSignPKCS1v15_v1(t *testing.T) {
	type args struct {
		data       []byte
		publicKey  *rsa.PublicKey
		privateKey *rsa.PrivateKey
	}
	publicKey, _ := LoadPublicKeyPem([]byte(PKCS1PublicKey))
	privateKey, _ := LoadPrivateKeyPem([]byte(PKCS1PrivateKey))

	tests := []struct {
		name          string
		args          args
		wantEncrypted string
		wantErr       bool
	}{
		{
			name: "Test rsa encrypt mode",
			args: args{
				data:       []byte("Hello RSA"),
				publicKey:  publicKey,
				privateKey: privateKey,
			},
			wantEncrypted: "g1W4EhpAm/pYU49N9rkbEeMsakkX73CjwGu1N0xFoyexQDON0Bm1dOy+yUaUwgVtvQexp1vVW0SS3d3AO/b8ag==",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := sha256.New()
			h.Write(tt.args.data)
			digest := h.Sum(nil)

			got, err := rsa.SignPKCS1v15(nil, rsaPrivateKey, crypto.SHA256, digest)
			gotEncrypted := strings.TrimSpace(encryptBase64(got))
			if err != nil {
				t.Errorf("SignPKCS1v15() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Logf("the SignPKCS1v15 got = \n%s", gotEncrypted)
			if gotEncrypted != tt.wantEncrypted {
				t.Errorf("SignPKCS1v15() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVerifyPKCS1v15_v1(t *testing.T) {
	type args struct {
		data       []byte
		sign       []byte
		publicKey  *rsa.PublicKey
		privateKey *rsa.PrivateKey
	}
	publicKey, _ := LoadPublicKeyPem([]byte(PKCS1PublicKey))
	privateKey, _ := LoadPrivateKeyPem([]byte(PKCS1PrivateKey))

	data, _ := decryptBase64("g1W4EhpAm/pYU49N9rkbEeMsakkX73CjwGu1N0xFoyexQDON0Bm1dOy+yUaUwgVtvQexp1vVW0SS3d3AO/b8ag==")

	tests := []struct {
		name          string
		args          args
		wantDecrypted string
		wantErr       bool
	}{
		{
			name: "Test rsa decrypt mode",
			args: args{
				sign:       data,
				data:       []byte("Hello RSA"),
				publicKey:  publicKey,
				privateKey: privateKey,
			},
			wantDecrypted: "",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := sha256.New()
			h.Write(tt.args.data)
			digest := h.Sum(nil)

			err := rsa.VerifyPKCS1v15(&rsaPrivateKey.PublicKey, crypto.SHA256, digest, tt.args.sign)
			if err != nil {
				t.Errorf("VerifyPKCS1v15() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
