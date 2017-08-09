package encoder

import (
	"encoding/base64"
)

func Base64(target string) string {
	return base64.StdEncoding.EncodeToString([]byte(target))
}
