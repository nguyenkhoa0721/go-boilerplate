package util

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

var shortSecretTable = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func HashSecret(secret string) string {
	hash := sha256.New()
	hash.Write([]byte(secret))
	return hex.EncodeToString(hash.Sum(nil))
}

func GenerateShortSecret(max int) (string, string, error) {
	secret := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, secret, max)
	if n != max {
		return "", "", err
	}

	for i := 0; i < len(secret); i++ {
		secret[i] = shortSecretTable[int(secret[i])%len(shortSecretTable)]
	}

	hash := sha256.New()
	hash.Write(secret)
	token := hash.Sum(nil)

	return string(secret), hex.EncodeToString(token), nil
}

func GenerateLongSecret(max int) (string, string, error) {
	secret := make([]byte, max)
	if _, err := rand.Read(secret); err != nil {
		return "", "", err
	}

	hash := sha256.New()
	hash.Write(secret)
	token := hash.Sum(nil)

	return hex.EncodeToString(secret), hex.EncodeToString(token), nil
}

func VerifyShortSecret(secret string, token string) bool {
	secretByte := []byte(secret)
	hash := sha256.New()
	hash.Write(secretByte)

	return hex.EncodeToString(hash.Sum(nil)) == token
}

func VerifyLongSecret(secret string, token string) bool {
	secretByte, err := hex.DecodeString(secret)
	if err != nil {
		return false
	}

	hash := sha256.New()
	hash.Write(secretByte)

	return hex.EncodeToString(hash.Sum(nil)) == token
}

func VerifyHMAC(secret, data, expect string) bool {
	hash := sha256.New()
	hash.Write([]byte(secret))
	hm := hmac.New(sha256.New, hash.Sum(nil))
	hm.Write([]byte(data))
	checkHash := hex.EncodeToString(hm.Sum(nil))
	return checkHash == expect
}

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
