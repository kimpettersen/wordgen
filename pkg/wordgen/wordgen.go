package wordgen

import (
	"time"

	"math/rand"
)

type Wordgen struct {
	Words []string
}

func (w Wordgen) Generate(count int, delimiter string) string {
	if count < 1 {
		count = 1
	}
	var words string
	for i := 1; i <= count; i++ {
		if words != "" {
			words += delimiter
		}
		words += w.random()
	}
	return words
}

func (w Wordgen) random() string {
	rand.Seed(time.Now().UnixNano())
	word := w.Words[rand.Intn(len(w.Words))]
	return word
}
