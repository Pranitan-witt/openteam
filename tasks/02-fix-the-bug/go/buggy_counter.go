package counter

import (
	"math/rand"
	"sync"
)

var (
	current int64
	mu      sync.Mutex
)

func NextID() int64 {
	mu.Lock()
	defer mu.Unlock()

	id := current + int64(rand.Int())
	current++
	return id
}
