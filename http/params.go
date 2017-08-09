package http

import (
	"bytes"
	"fmt"
	"sync"
)

var bufPool = sync.Pool{
	New: func() interface{} { return bytes.Buffer{} },
}

type Params []Param

type Param struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func NewParams() *Params {
	return &Params{}
}

func (p *Params) Add(key, value string) *Params {
	(*p) = append((*p), Param{key, value})
	return p
}

func (p *Params) AddInt64(key string, value int64) *Params {
	(*p) = append((*p), Param{key, fmt.Sprintf("%d", value)})
	return p
}

func (p *Params) Set(key, value string) *Params {
	for idx, param := range *p {
		if param.Name == key {
			(*p)[idx].Value = value
		}
	}
	return p
}

func (p *Params) Get(key string) (bool, string) {
	for _, param := range *p {
		if param.Name == key {
			return true, param.Value
		}
	}

	return false, ""
}

func (p *Params) Rename(old, new string) {
	for id, param := range *p {
		if param.Name == old {
			(*p)[id].Name = new
			return
		}
	}
}

func (p *Params) String() (ret string) {
	switch p.Len() {
	case 0:
	case 1:
		ret = fmt.Sprintf("%s=%s", (*p)[0].Name, (*p)[0].Value)
	default:
		buf := bufPool.Get().(bytes.Buffer)
		defer bufPool.Put(buf)

		for _, param := range *p {
			buf.WriteString(param.Name)
			buf.WriteByte('=')
			buf.WriteString(param.Value)
			buf.WriteByte('&')
		}

		ret = string(buf.Bytes()[:buf.Len()-1])
	}

	return ret
}

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

func (p *Params) Len() int {
	slice := ([]Param)(*p)
	return len(slice)
}

func (p *Params) Less(i, j int) bool {
	slice := ([]Param)(*p)
	return slice[i].Name < slice[j].Name
}

func (p *Params) Swap(i, j int) {
	slice := ([]Param)(*p)
	slice[i], slice[j] = slice[j], slice[i]
}
