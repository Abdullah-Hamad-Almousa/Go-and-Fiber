package main

import (
	"flag"
	"fmt"
	"strings"
)

func CountStats(text string) (int, int) {

	words := len(strings.Fields(text))
	lines := len(strings.Split(strings.TrimSpace(text), "\n"))

	if text == "" {
		lines = 0
	}
	return words, lines

}

func main() {

	input := flag.String("text", "", "Text to analyze")
	flag.Parse()

	words, lines := CountStats(*input)
	fmt.Printf("Words: %d | Lines: %d\n", words, lines)

}
