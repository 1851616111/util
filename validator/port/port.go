package validator

import (
	"github.com/1851616111/util/validator"
	"github.com/pkg/errors"
	"regexp"
)

var Err_Format_Invalid error = errors.New("port format invalid")

var portRegexp *regexp.Regexp = regexp.MustCompile(`^([1-9]|[1-9]\d{1,3}|[1-5]\d{4}|6[0-5]{2}[0-3][0-5])$`)

func Validate(port string) error {
	return validator.Validate(portRegexp, port)
}

func ValidateInt(port int) error {
	if port < 0 || port > 65535 {
		return Err_Format_Invalid
	}

	return nil
}
