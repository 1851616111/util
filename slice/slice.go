package util

import (
	"errors"
)

func IndexAddString(s *[]string, idx int, str string) error {
	if s == nil {
		return errors.New("util.slice.IndexAddString: nil slice")
	}

	if idx > len(*s) || idx < 0 {
		return errors.New("util.slice.IndexAddString: invalidated index")
	}

	s_copy := append(*s, "") //扩容s及copy
	copy(s_copy[idx+1:], s_copy[idx:len(s_copy)-1])
	s_copy[idx] = str
	*s = s_copy

	return nil
}

func IndexRemoveString(s *[]string, idx int) error {
	if s == nil {
		return errors.New("util.slice.IndexRemoveString: nil slice")
	}

	if idx >= len(*s) || idx < 0 {
		return errors.New("util.slice.IndexRemoveString: invalidated index ")
	}

	s_copy := *s
	copy(s_copy[idx:], s_copy[idx+1:])
	s_copy = s_copy[:len(s_copy)-1] //缩减slice大小
	*s = s_copy

	return nil
}
