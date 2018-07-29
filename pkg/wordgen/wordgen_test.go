package wordgen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateWord(t *testing.T) {
	generator := Wordgen{
		Words: []string{
			"one",
			"two",
			"three",
		},
	}
	words := generator.Generate(2, "-")
	assert.NotEmpty(t, words)
	assert.Contains(t, words, "-")
}

func TestDefaultValues(t *testing.T) {
	generator := Wordgen{
		Words: []string{
			"one",
			"two",
			"three",
		},
	}
	s := generator.Generate(0, ".")
	words := strings.Split(s, ".")

	// Default is 1. Assert that there are no delimiters.
	// i.e strings.Split returns one result
	assert.Len(t, words, 1)

	// And that the string contains at least a three letter wor
	assert.True(t, len(s) >= 3)
}
