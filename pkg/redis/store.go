package redis

import (
	"time"

	"github.com/go-redis/redis"
)

type Store struct {
	client     *redis.Client
	expiration time.Duration
}

func NewStore(client *redis.Client, expiration time.Duration) *Store {
	return &Store{client: client, expiration: expiration}
}

func (s *Store) Get(key string) ([]byte, bool) {
	if value, err := s.client.Get(key).Result(); err == nil {
		return []byte(value), true
	}
	return nil, false
}

func (s *Store) Set(key string, value []byte) {
	s.client.Set(key, value, s.expiration)
}
