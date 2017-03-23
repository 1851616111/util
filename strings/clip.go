package strings

import "strings"

var DB_ARRAY_MID string = `^,^`
var DB_ARRAY_PERFIX string = `[^[`
var DB_ARRAY_SUFFIX string = `]^]`

var DB_OBJECT_MID string = `^:^`
var DB_OBJECT_PERFIX string = `{^{`
var DB_OBJECT_SUFFIX string = `}^}`

func Clip(s *string, left, mid, right string) []string {
	if s == nil {
		return  nil
	}

	result := strings.TrimSpace(*s)

	if left != "" {
		result = strings.TrimLeft(result, left)
	}

	if right != "" {
		result = strings.TrimRight(result, right)
	}

	return strings.Split(result, mid)
}

func ClipDBArray(s *string) []string {
	if s == nil {
		return nil
	}

	return Clip(s, DB_ARRAY_PERFIX, DB_ARRAY_MID, DB_ARRAY_SUFFIX)
}

func ClipDBObject(s *string) []string {
	if s == nil {
		return nil
	}
	return Clip(s, DB_OBJECT_PERFIX, DB_OBJECT_MID, DB_OBJECT_SUFFIX)
}
