package text

import (
	"math/rand"
	"time"
)

func RandomNumber(i int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(i)
}
