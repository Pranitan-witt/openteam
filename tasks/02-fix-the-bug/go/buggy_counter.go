package counter

import (
	"math/rand"
	"time"
)

var current int64

func NextID() int64 {
	id := current + int64(rand.Int())
	time.Sleep(0)
	current++
	return id
}
