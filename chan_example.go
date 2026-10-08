package main

import "fmt"

// Goal: Create a typed communication pipeline to safely transfer
// data and close it when done.

func RunChanExample() {
	pipeline := make(chan int, 3)

	pipeline <- 10
	pipeline <- 20
	pipeline <- 30
	close(pipeline)

	for item := range pipeline {
		fmt.Printf("[chan] Received from pipeline: %d\n", item)
	}

}
