package text

import (
	"math/rand"
	"time"
)

func Random(texts []Text) Text {
	text := texts[RandomNumber(len(texts))]
	return text
}

func RandomNumber(i int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(i)
}
