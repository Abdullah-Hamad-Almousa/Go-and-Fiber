package main

import "fmt"

// Goal: Enforce directional type safety using arrow operators for
// sending (chan<-) and receiving (<-chan).

func produce(out chan<- string) {
	out <- "package_ready"
}

func consume(in <-chan string) {
	data := <-in
	fmt.Printf("[<-] Processed incoming item: %s\n", data)
}

func RunArrowExample() {
	ch := make(chan string, 1)
	produce(ch)
	consume(ch)
}
