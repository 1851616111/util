package validator

import (
	"github.com/1851616111/util/validator"
	"github.com/pkg/errors"
	"regexp"
)

var Err_Format_Invalid error = errors.New("ip format invalid")

var ipRegexp *regexp.Regexp = regexp.MustCompile(`^(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[0-9]{1,2})(\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[0-9]{1,2})){3}$`)

func Validate(ip string) error {
	return validator.Validate(ipRegexp, ip)
}
