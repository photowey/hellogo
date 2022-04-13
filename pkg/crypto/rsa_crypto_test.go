package crypto

import (
	"crypto/rsa"
	"testing"
)

// openssl genrsa -out rsa_private_key.pem 2048
// openssl pkcs8 -topk8 -inform PEM -in rsa_private_key.pem -outform PEM -nocrypt -out pkcs8_private_key.pem
// openssl rsa -in pkcs8_private_key.pem -pubout -out pkcs8_public_key.pem

// 注意: 此处-多行字符串 - 不要带空格
// 错误示例:
// 空格空格空格空格MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDRdMvcqBqIcS2K
// 空格空格空格空格ZmHXdFrkhv97NU7OuQZNKSSmN7SzYp5jhrHZ+30aU82pK3s0Wh75SK3V3iivsSqK

var (
	PRIVATE_KEY_PEM = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDRdMvcqBqIcS2K
ZmHXdFrkhv97NU7OuQZNKSSmN7SzYp5jhrHZ+30aU82pK3s0Wh75SK3V3iivsSqK
Xdn/VJIsMsTcKpsNsbwouqLWU24Rt6XlRN7ypE0sq9StfE2NrxhGrYSoHwz4Fey3
2lUDibBV8Jz/ndEczCeMqVa0WGQZaRTDJ0+p0f+XHlUKnm12PpwOU28RBKJEtf+u
nT1O/G3tw84QKKvpWm6YinBfLaIN6CvpfCgwoEwinaW38sF5n+WF+1yJg5RHz3fj
oLylkZ1kq+HTSJbAxdOlTxE6P9uOBWuQstfP5B1oyol21L0Q5SngE1w+GuyBsym3
VHfZjKnBAgMBAAECggEAJrTtw34xQQuhPdVDuXwgG+Eyr0MfqCYbGwFCa5EZzJip
6nyGu2C1MrtP4zZM01TgKnMa5M3kOOuFkAJd+chYJuDO7lzVQIea3O/4jaDnU1Db
0G2UwcRJGrs6V0EEV+2Gj/Ea9bKfQ+1RTHyFf1zfgFIxwS8Z6Ld1i5HzmupGvUcP
IMBq51zsOOgWjXoYmVypW37KRkO2PrbGTv++j9xsniIwEsFa9T04nBtr55/bwG6A
482FSGgA/wGoYOBkLgLGrHyhyb5rbYmrmDYjlRGxAsgW5M10NnNTdH4/rCVEoKZy
ARgo+T4Yx5/aG0eoiVYnxgBxtmSuK6BfbbzBwxCtAQKBgQD56EuZt/+Ttzo1MZXq
kWkWPQKR7geLUA9eLvjDOyTN+LhYybobuayhOdQFd94jkpRDBgxciNiW4ZCqBQ7Y
o/e2NACFK1sMlW0oBJoHjqoVTEf4zPrZE3NXEZ+DalNERTP5RWtQVeWiQLGDWXu5
lDV/2Wb0CKSrUzWoYbT5wcMkcQKBgQDWkAo+CNLU7t8Kp9/G8qyACzAluQZr++wB
Mbo3P4Nm1MATiPxcmPEApLL+NkPnEMRFggD9RHKOceLpK7gGtwd7EpTC0FyL7N7K
LVOGyxPK4XjdW9y3fFWjXlyAADmvclo3I5mrVsF/fusjrwNozPlej46sFfVSvZpW
BVEKInlCUQKBgFYgIn3VlbUYvmy6rAntFyRc3lfgiqUOcF1rHrxZ07NQdmfTDdqk
QMRHo7RGvT3RNFts6DGzz8Ef97VAjP9IcwROI6fBa12I2deizlqwZhl08pDiQ8kw
FN366thlelDa8LCFphhHoh4q4YJ9AlDQrSdki0CrAeUxhhqDodU5W8FhAoGACnn+
BgloQf1n5p1Lz+sJgTcKoszmA1/JpkVHUmtHfz3kIxANQng1TFK+aqs7pYysCLjQ
U42ECnCFI91+ntVRCDFgEiMaYbqvJjAQZASGcuGPhT93B36Kr+9rlNRt/KMYxJh4
aGVUDWWlFZj2EafwhS19SenIkMGxR4H68wO76AECgYEA26q+0IydBGFJHm09/VVr
xNsYZGuNVKDiXeEGRi7LzrFykLG4NdPF+owCqE11fVj6hu93Qi8BMCJOXreHKSOg
i3Al02XDQML5zNyZ/0/eh90QsdhzoeRGe4q9XYahntizWA7cgHwIkLsvSgmkwfu1
aD9gjyBhzRYWn/NcA+n4BnE=
-----END PRIVATE KEY-----`

	PUBLIC_KEY_PEM = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0XTL3KgaiHEtimZh13Ra
5Ib/ezVOzrkGTSkkpje0s2KeY4ax2ft9GlPNqSt7NFoe+Uit1d4or7Eqil3Z/1SS
LDLE3CqbDbG8KLqi1lNuEbel5UTe8qRNLKvUrXxNja8YRq2EqB8M+BXst9pVA4mw
VfCc/53RHMwnjKlWtFhkGWkUwydPqdH/lx5VCp5tdj6cDlNvEQSiRLX/rp09Tvxt
7cPOECir6VpumIpwXy2iDegr6XwoMKBMIp2lt/LBeZ/lhftciYOUR89346C8pZGd
ZKvh00iWwMXTpU8ROj/bjgVrkLLXz+QdaMqJdtS9EOUp4BNcPhrsgbMpt1R32Yyp
wQIDAQAB
-----END PUBLIC KEY-----`
)

var (
	PRIVATE_KEY = `MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDwkG8hAYkYtaJCJPJZs8nAUJKJtyucuX7AsHr8HiUm/qHdB/z0VO+46Sgpy494v8zDqkf++R/gDIKwKIgxBRISdcEE9KL7z9uf8MH1LaXAGyQbOvHPpjTsHoxvd5pwlzmdl6V8yTyc//nD2Zd1n7AIB9vEz7/0gnK4U6LlcG4rPQ+AXOHJT2Ye0KEkqRYEeFAMPAJ0I5NbxTbbmHxZsWRfg5u13zEvZUok5gn1fhBMGgvIzoFI+tVCmuFOyHzzcVN+GWdB8qix6KJ0pFndpjO75+pSZ1rbzA8q7cTc2XOqvfRRuTzEVUnkn1AAxpQxWWQLt7+1y7NYq1QpVfyJQY8XAgMBAAECggEAMY3PFoF2eDySHAX4VjoOySrPxMaSmUosrgysiNq81HeHaDyJAtusNe7xcwqXl0wLJhMhmYg/5KsvEJKI23Ar9NmE1Tx/hlee6idyDMtGZXgxy9osjZWzOSBh1WRRz/eWyru+PyJE9iJUlN9xr60hWcYkozoYVIm6j/XKjHDIgOdA1jdiTt9Tee4jdrIE3UlJnz5nOC0ObNYQyn74otZUDD4Q1Akc5VFw71v43yAit3YfRfDzTffRiOsHFU97Zm3P91K28zaL01/Udly4452z9l6RQn2DuoMTBeOa6bW2zaryZwWXhrLxLFGZnuT0Y4Y7fq9g6DIC5sWY1+aMSAZ64QKBgQD7tkdsfQcP1U52rEhRWGgi8y3+gNXPjC85452k6u5ic+rxKIbGXfGiENUJ7eJSB1VaRQ4GJGDHckVC0m86H0pp3TJjRAcfcTtWQNGEkGOGXFcLfyv3XQJZO/R/JiZA58ki/tnJXUeQDASi2PoOegihwPyHAiGqfQBt6DKNEf5P2wKBgQD0qYoIVPl7ga8enxw03sRrvc3zI2N8K7yQ3n+F242NlzWB7nmL/sRR+Kqr7VgzPWURd07U0gQHRs3ri3hsNSS+S6uBL81dhosEBpfpF7xaAsBI6/uhiqS94THXtBvwoWWIPbXV7bg+TCZb1Tkn4NUjl6qeH7KffCkfzC+4T4owdQKBgHXwM0kGz+9AP7JlCMFsjeaKUKulzAiiI2KZXrzuAuIIdDWoinWZgOUxw9ASMO/EVd6k6mxAUURCK2ei19DA5onxxXEzZ674zBee6UyWprtGutY9MOmHH9mVuPp2cogI0npeNMcfuK8qomV3CrsLvM2lce0EN6f9R903ZbAplSHvAoGBAOGymLQw92jY7FxaOSIIsrgVgyydY8QDcNNPfeJU9FM+v0yaWqQrpJ1dwMwTij7Sjy66mg03pCG2TnXQ1okzyS6HDbKCD3gMiqdvEaokjEuYqBtdHmcaR/kmCd3Lt86WhNxW2Sx9PG6zvtsOash/3v8ThYCTCyuW1TURHqBvN7KVAoGBAJga1i0REzxO39K853INjfEeuzqZ7gkiEb0/wPVpVsUQcQLCO9fTYiC08DoRFa+YUqO2NU4OwT0wBKABkz75X5avXYCFD1STquNd9Csr9Rt/wcb4riLiarnm+2UCt6cFg/ONAIrbt0lrINCW+acvkWTOZv8Op9qORDOU6vZ0fObs`
	PUBLIC_KEY  = `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA8JBvIQGJGLWiQiTyWbPJwFCSibcrnLl+wLB6/B4lJv6h3Qf89FTvuOkoKcuPeL/Mw6pH/vkf4AyCsCiIMQUSEnXBBPSi+8/bn/DB9S2lwBskGzrxz6Y07B6Mb3eacJc5nZelfMk8nP/5w9mXdZ+wCAfbxM+/9IJyuFOi5XBuKz0PgFzhyU9mHtChJKkWBHhQDDwCdCOTW8U225h8WbFkX4Obtd8xL2VKJOYJ9X4QTBoLyM6BSPrVQprhTsh883FTfhlnQfKoseiidKRZ3aYzu+fqUmda28wPKu3E3Nlzqr30Ubk8xFVJ5J9QAMaUMVlkC7e/tcuzWKtUKVX8iUGPFwIDAQAB`
)

func TestEncryptOAEP(t *testing.T) {
	type args struct {
		data       []byte
		publicKey  *rsa.PublicKey
		privateKey *rsa.PrivateKey
	}
	publicKey, _ := LoadPublicKeyPem([]byte(PUBLIC_KEY_PEM))
	privateKey, _ := LoadPrivateKeyPem([]byte(PRIVATE_KEY_PEM))

	tests := []struct {
		name          string
		args          args
		wantEncrypted string
		wantDecrypted string
		wantErr       bool
	}{
		{
			name: "Test rsa Codec OAEP mode",
			args: args{
				data:       []byte("Hello RSA"),
				publicKey:  publicKey,
				privateKey: privateKey,
			},
			wantDecrypted: "Hello RSA",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEncrypted, err := EncryptByPublicKeyOAEP(tt.args.data, tt.args.publicKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptByPublicKeyOAEP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotDecrypted, err := DecryptByPrivateKeyOAEP(gotEncrypted, tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptByPrivateKeyOAEP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotDecrypted != tt.wantDecrypted {
				t.Errorf("EncryptByPublicKeyOAEP() And DecryptByPrivateKeyOAEP() gotDecrypted = %v, want %v", gotEncrypted, tt.wantDecrypted)
			}
		})
	}
}

func TestEncryptBlockOAEP(t *testing.T) {
	type args struct {
		data       []byte
		publicKey  *rsa.PublicKey
		privateKey *rsa.PrivateKey
	}
	publicKey, _ := LoadPublicKeyPem([]byte(PUBLIC_KEY_PEM))
	privateKey, _ := LoadPrivateKeyPem([]byte(PRIVATE_KEY_PEM))

	tests := []struct {
		name          string
		args          args
		wantEncrypted string
		wantDecrypted string
		wantErr       bool
	}{
		{
			name: "Test rsa block Codec OAEP mode",
			args: args{
				data:       []byte("Hello RSA"),
				publicKey:  publicKey,
				privateKey: privateKey,
			},
			wantDecrypted: "Hello RSA",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEncrypted, err := EncryptByPublicKeyBlockOAEP(tt.args.data, tt.args.publicKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptByPublicKeyOAEP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotDecrypted, err := DecryptByPrivateKeyBlockOAEP(gotEncrypted, tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptByPrivateKeyBlockOAEP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDecrypted != tt.wantDecrypted {
				t.Errorf("EncryptByPublicKeyBlockOAEP() And DecryptByPrivateKeyBlockOAEP() gotDecrypted = %v, want %v", gotDecrypted, tt.wantDecrypted)
			}
		})
	}
}

func TestEncryptPKCS1(t *testing.T) {
	type args struct {
		data       []byte
		publicKey  *rsa.PublicKey
		privateKey *rsa.PrivateKey
	}
	publicKey, _ := LoadPublicKeyPem([]byte(PUBLIC_KEY_PEM))
	privateKey, _ := LoadPrivateKeyPem([]byte(PRIVATE_KEY_PEM))

	tests := []struct {
		name          string
		args          args
		wantEncrypted string
		wantDecrypted string
		wantErr       bool
	}{
		{
			name: "Test rsa Codec PKCS1 mode",
			args: args{
				data:       []byte("Hello RSA"),
				publicKey:  publicKey,
				privateKey: privateKey,
			},
			wantDecrypted: "Hello RSA",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEncrypted, err := EncryptByPublicKeyPKCS1(tt.args.data, tt.args.publicKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptByPublicKeyPKCS1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotDecrypted, err := DecryptByPrivateKeyPKCS1(gotEncrypted, tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptByPrivateKeyPKCS1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotDecrypted != tt.wantDecrypted {
				t.Errorf("EncryptByPublicKeyPKCS1() And DecryptByPrivateKeyPKCS1() gotDecrypted = %v, want %v", gotEncrypted, tt.wantDecrypted)
			}
		})
	}
}

func TestEncryptBlockPKCS1(t *testing.T) {
	type args struct {
		data       []byte
		publicKey  *rsa.PublicKey
		privateKey *rsa.PrivateKey
	}
	publicKey, _ := LoadPublicKeyPem([]byte(PUBLIC_KEY_PEM))
	privateKey, _ := LoadPrivateKeyPem([]byte(PRIVATE_KEY_PEM))

	tests := []struct {
		name          string
		args          args
		wantEncrypted string
		wantDecrypted string
		wantErr       bool
	}{
		{
			name: "Test rsa block Codec PKCS1 mode",
			args: args{
				data:       []byte("Hello RSA"),
				publicKey:  publicKey,
				privateKey: privateKey,
			},
			wantDecrypted: "Hello RSA",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEncrypted, err := EncryptByPublicKeyBlockPKCS1(tt.args.data, tt.args.publicKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptByPublicKeyBlockPKCS1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotDecrypted, err := DecryptByPrivateKeyBlockPKCS1(gotEncrypted, tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptByPrivateKeyBlockPKCS1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDecrypted != tt.wantDecrypted {
				t.Errorf("EncryptByPublicKeyBlockPKCS1() And DecryptByPrivateKeyBlockPKCS1() gotDecrypted = %v, want %v", gotDecrypted, tt.wantDecrypted)
			}
		})
	}
}
