package strings

import "strings"

func Clip(s string, left, mid, right string) []string {
	if left != "" {
		s = strings.TrimLeft(s, left)
	}

	if right != "" {
		s = strings.TrimRight(s, right)
	}

	s = strings.TrimSpace(s)
	return strings.Split(s, mid)
}
