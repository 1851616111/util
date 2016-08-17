package main

import (
	"fmt"
	"sync"
)

type generator interface {
	gen() string
}

type safeID struct {
	idStarting uint64 //atomic
	sync.Mutex
}

func (i *safeID) gen() string {
	return fmt.Sprintf("%d", i.inc())
}

func (i *safeID) inc() uint64 {
	var ret uint64
	i.Lock()
	i.idStarting += 1
	ret = i.idStarting
	i.Unlock()

	return ret
}
