package text

import (
  // "fmt"
  "math/rand"
  "time"
)

type Text struct {
  ja string
  en string
}

func Random() Text {
  texts := []Text{
    {"最高!", "Great!"},
    {"強す", "Strong!"},
    {"す!", "Strong!"},
    {"すぎ!", "Strong!"},
  }

  return random(texts)
}

func random(texts []Text) Text {
  text := texts[randomNumber(len(texts))]
  return text
}

func randomNumber(i int) int {
  rand.Seed(time.Now().UnixNano())
  return rand.Intn(i)
}
