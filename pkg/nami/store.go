package nami

import (
	"fmt"
)

type Store interface {
	Get(key string) ([]byte, bool)
	Set(key string, value []byte)
}

func generateDirectionKey(origin, destination string) string {
	return fmt.Sprintf("directions:%s:%s", origin, destination)
}

func generateStaticMapKey(origin, destination string) string {
	return fmt.Sprintf("staticmap:%s:%s", origin, destination)
}
