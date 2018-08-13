package redis

import (
	"encoding/json"
	"log"
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

func (s *Store) Get(key string) (interface{}, bool) {
	if value, err := s.client.Get(key).Result(); err == nil {
		var result interface{}
		json.Unmarshal([]byte(value), &result)
		return result, true
	}
	return nil, false
}

func (s *Store) Set(key string, value interface{}) {
	if j, err := json.Marshal(value); err == nil {
		s.client.Set(key, j, s.expiration)
	} else {
		log.Print(err)
	}
}
