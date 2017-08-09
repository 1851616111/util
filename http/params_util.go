package http

import (
	"bytes"
	"sort"
)

func (ps *Params) UrlString() string {
	buf := bufPool.Get().(bytes.Buffer)
	defer bufPool.Put(buf)

	sort.Sort(ps)
	for _, param := range *ps {
		buf.WriteString(param.Name)
		buf.WriteRune('=')
		buf.WriteString(param.Value)
		buf.WriteRune('&')
	}

	var ret string = buf.String()
	if ps.Len() > 0 {
		ret = ret[:len(ret)-1]
	}

	return ret
}
