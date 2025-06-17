package counter

import (
	"math/rand"
	"sync/atomic"
)

var current int64

func NextID() int64 {
	id := atomic.AddInt64(&current, 1) + int64(rand.Int())
	current++
	return id
}
