package encoder

type Encoder interface {
	Method() string
	Encode(text, key string) string
}

var Hmac_Sha_1 encoder = "HmacSHA1"
var Hmac_Sha_256 encoder = "HmacSHA256"
var Base_64 encoder = "Base_64"

type encoder string

func (s *encoder) Method() string {
	return string(*s)
}

func (s *encoder) Encode(text, key string) string {
	switch *s {
	case Hmac_Sha_1:
		return HmacSha1(text, key)
	case Hmac_Sha_256:
		return HmacSha256(text, key)
	case Base_64:
		return Base64(text)
	default:
		return ""
	}
}

type other struct {
	method   string
	EncodeFn func(text, key string) string
}

func (e other) Method() string {
	return e.method
}

func (e other) Encode(text, key string) string {
	if e.EncodeFn != nil {
		return e.EncodeFn(text, key)
	}
	return ""
}

func NewEncoder(m string, fn func(string, string) string) Encoder {
	return other{
		method:   m,
		EncodeFn: fn,
	}
}
