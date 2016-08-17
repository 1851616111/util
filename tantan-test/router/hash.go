package router

import (
	"errors"
	"github.com/stathat/consistent"
)

// this interface is used for other hash algorithms
// now I use consistent hash
type hasher interface {
	hash(string) (string, error)
}

// as this is a test rehash json isn't done
type consistentHash struct {
	*consistent.Consistent
}

func (c *consistentHash) hash(key string) (string, error) {
	return c.Get(key)
}

type hash struct {
	hasher
}

func (b *hash) Route(key string) (string, error) {
	return b.hasher.hash(key)
}

func NewHashRouter(nodes []string) (Router, error) {
	if len(nodes) == 0 {
		return nil, errors.New("no nodes to route by hash")
	}

	c := consistent.New()
	c.NumberOfReplicas = 10000 //to make more equal

	c.Set(nodes)

	hasher := &consistentHash{c}
	return &hash{hasher}, nil
}
