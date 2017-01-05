package redis

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var cacher Cache

type Cache interface {
	Get(key string) ([]byte, error)
	Set(key string, data []byte) error
}

func CreatePool(server, auth string) Cache {
	return &cache{pool: &redis.Pool{
		MaxIdle:     10,
		MaxActive:   10,
		Wait:        true,
		IdleTimeout: 4 * time.Minute,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if auth != "" {
				if _, err := c.Do("AUTH", auth); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}}
}

type cache struct {
	pool *redis.Pool
}

func (p *cache) Get(key string) ([]byte, error) {
	c := p.pool.Get()
	defer c.Close()
	return redis.Bytes(c.Do("GET", key))
}

func (p *cache) Set(key string, data []byte) (err error) {
	c := p.pool.Get()
	defer c.Close()
	_, err = redis.Bytes(c.Do("SET", key, data))

	return
}
