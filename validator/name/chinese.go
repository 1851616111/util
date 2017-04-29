package name

import (
	"errors"
	"regexp"

	"github.com/1851616111/util/validator"
)

var Err_Format_Invalid error = errors.New("chinese name format invalid")

var nameRegexp *regexp.Regexp = regexp.MustCompile(`^[\p{Han}]{2,5}$`)

func Validate(name string) error {
	return validator.Validate(nameRegexp, name)
}
