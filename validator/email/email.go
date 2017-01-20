package email

import (
	"github.com/1851616111/util/validator"
	"github.com/pkg/errors"
	"regexp"
)

var Err_Format_Invalid error = errors.New("email format invalid")
var emailRegexp *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)

func Validate(email string) error {
	return validator.Validate(emailRegexp, email)
}