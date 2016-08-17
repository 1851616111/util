package router

import (
	"errors"
	"strconv"
	"sync"
)

type block struct {
	begin        uint64
	size         uint64
	max          uint64
	routes       []string
	compatible   bool
	sync.RWMutex //used for concurrency and dynamically extend nodes
}

func (b *block) Route(key string) (string, error) {
	b.RLock()
	defer b.RUnlock()

	k_uint64, err := strconv.ParseUint(key, 10, 64)
	if err != nil {
		return "", err
	}

	if k_uint64 < b.begin {
		return "", errors.New("route key is invalid")
	}

	if !b.compatible && k_uint64 >= b.max {
		return "", errors.New("route key is out of processable limit")
	}

	blockIndex := int((k_uint64 - b.begin) / b.size)
	if b.compatible && blockIndex >= len(b.routes) {
		blockIndex = len(b.routes) - 1
	}

	return b.routes[blockIndex], nil
}

func NewBlockRouter(size, begin uint64, nodes []string, compatible bool) (Router, error) {
	if len(nodes) == 0 {
		return nil, errors.New("no nodes to route by block")
	}

	if size == 0 {
		return nil, errors.New("missing param size")
	}

	if begin == 0 {
		return nil, errors.New("missing begin size")
	}

	return &block{
		begin:      begin,
		size:       size,
		max:        begin + uint64(len(nodes))*size,
		routes:     nodes,
		compatible: compatible,
	}, nil
}
