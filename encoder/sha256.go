package encoder

import (
	"crypto/hmac"
	"crypto/sha256"
)

func HmacSha256(text, key string) string {
	hmacSha256 := hmac.New(sha256.New, []byte(key))
	hmacSha256.Write([]byte(text))

	return string(hmacSha256.Sum(nil))
}
