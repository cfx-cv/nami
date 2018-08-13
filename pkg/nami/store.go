package nami

import (
	"fmt"
)

type Store interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

func generateKey(origin, destination string) string {
	return fmt.Sprintf("%s:%s", origin, destination)
}
