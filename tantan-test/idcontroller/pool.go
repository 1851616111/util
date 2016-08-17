package main

import "time"

func NewPool(g generator, o *option) IDPool {
	if o == nil {
		o = defaultPoolOption
	}

	p := &pool{
		poolCh: make(chan string, o.poolSize),
	}

	go func() {
		for {
			select {
			case p.poolCh <- g.gen():
			}
		}
	}()

	return p
}

var defaultPoolOption = &option{
	poolSize: 20,
	period:   time.Second,
}

type option struct {
	poolSize uint32
	period   time.Duration
}

type IDPool interface {
	GenID() string
}

type pool struct {
	poolCh chan string
}

func (p *pool) GenID() string {
	return <-p.poolCh
}
