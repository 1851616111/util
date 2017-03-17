package strings

import "strings"

func Clip(s string, left, mid, right string) []string {
	s = strings.TrimSpace(s)

	s = strings.TrimLeft(s, left)
	s = strings.TrimRight(s, right)

	return strings.Split(s, mid)
}
