package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/kimpettersen/wordgen/pkg/wordgen"
)

func main() {
	count := flag.Int("count", 4, "The amount of words you want. Must be >= 1")
	delimiter := flag.String("delimiter", ".", "The delimiter used. Example: \".\" means word1.word2")
	path := flag.String("path", "", "path to file that contains words. One word per line")
	flag.Parse()

	f, err := ioutil.ReadFile(*path)
	if err != nil {
		fmt.Printf("Could not read file: %+v", err)
	}
	words := strings.Split(string(f), "\n")
	w := wordgen.Wordgen{
		Words: words,
	}

	fmt.Println(w.Generate(*count, *delimiter))
}
