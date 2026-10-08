package main

import (
	"fmt"
	"time"
)

// Goal: Launch an independent, non-blocking background task
// without stalling the main execution flow.

func backgroundWorker(taskName string) {

	fmt.Printf("[go] String background task: %s\n", taskName)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("[go] finished background task: %s\n", taskName)

}

func RunGoExample() {
	go backgroundWorker("SendEmailNotification")
	fmt.Println("[go] Main thread continues without waiting...")
	time.Sleep(150 * time.Millisecond)
}
