package main

import (
	"fmt"
	"sync"
)

// Goal: Coordinate multiple concurrent workers so the caller
// waits until all finish their jobs completely.

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[sync] Worker %d finished work\n ", id)
}

func RunSyncExample() {
	var wg sync.WaitGroup

	for i := range 3 {
		i++
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("[sync] All workers completed!")
}
