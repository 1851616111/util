package encoder

import (
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	encoder := Base64("NSI3UqqD99b/UJb4tbG/xZpRW64=")
	fmt.Println(encoder)

}
