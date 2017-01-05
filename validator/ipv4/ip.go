package validator

import (
	"github.com/pkg/errors"
	"regexp"
)

var Err_Format_Invalid error = errors.New("ip format invalid")
var ipRegxp *regexp.Regexp = regexp.MustCompile(`^(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[0-9]{1,2})(\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[0-9]{1,2})){3}$`)

func Validate(ip string) error {
	if ip == "" {
		return Err_Format_Invalid
	}

	if !ipRegxp.MatchString(ip) {
		return Err_Format_Invalid
	}

	return nil
}
