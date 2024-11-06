package apc

import (
	"testing"
)

func TestGet(t *testing.T) {
	key := "hello"
	value := "golang"
	Set(key, value, -1)

	v, exists := Get(key)
	if !exists || value != v.(string) {
		t.Fatal("failed to get: hello golang")
	}
}
