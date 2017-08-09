package encoder

import (
	"crypto/hmac"
	"crypto/sha1"
)

func HmacSha1(text, key string) string {
	hmacSha1 := hmac.New(sha1.New, []byte(key))
	hmacSha1.Write([]byte(text))

	return string(hmacSha1.Sum(nil))
}
