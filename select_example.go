package main

import (
	"fmt"
	"time"
)

// Goal: Listen to multiple channel operations simultaneously
// and execute whichever is ready first (timeout / racing).

func RunSelectExample() {
	fastCh := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		fastCh <- "Fast server response"
	}()

	select {
	case res := <-fastCh:
		fmt.Printf("[select] Handled: %s\n", res)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("[select] Request timeout out")
	}

}
