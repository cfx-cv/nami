package nami

import (
	"fmt"
)

type Store interface {
	Get(key string) ([]byte, bool)
	Set(key string, value []byte)
}

func generateKey(origin, destination string) string {
	return fmt.Sprintf("%s:%s", origin, destination)
}
