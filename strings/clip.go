package strings

import "strings"

var DB_ARRAY_MID string = `^,^`
var DB_ARRAY_PERFIX string = `[^[`
var DB_ARRAY_SUFFIX string = `]^]`

var DB_OBJECT_MID string = `^,^`
var DB_OBJECT_PERFIX string = `{^{`
var DB_OBJECT_SUFFIX string = `}^}`

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

func ClipDBArray(s string) []string {
	return Clip(s, DB_ARRAY_PERFIX, DB_ARRAY_MID, DB_ARRAY_SUFFIX)
}

func ClipDBObject(s string) []string {
	return Clip(s, DB_OBJECT_PERFIX, DB_OBJECT_MID, DB_OBJECT_SUFFIX)
}
