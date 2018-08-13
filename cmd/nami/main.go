package main

import (
	"os"
	"strconv"
	"time"

	"github.com/cfx-cv/nami/pkg/server"
	"github.com/go-redis/redis"
)

func main() {
	client := newClient()
	expiration := parseExpiration()

	server.NewServer(client, expiration).Start()
}

func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})

	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}
	return client
}

func parseExpiration() time.Duration {
	value, _ := strconv.Atoi(os.Getenv("REDIS_EXPIRATION"))
	return time.Duration(value)
}
