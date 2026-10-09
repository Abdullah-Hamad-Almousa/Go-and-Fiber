package main

import "testing"

func TestCountStats(t *testing.T) {

	input := "Hello Go\nWelcome to Week 1"
	words, lines := CountStats(input)

	if words != 5 || lines != 2 {
		t.Errorf("Expected 5 words, 2 lines; got %d words, %d lines",
			words, lines)
	}

}
