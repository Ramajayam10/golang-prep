package main

import (
	"fmt"
	"sync"
	"time"
)

var fetchMu sync.Mutex
var lastFetch time.Time

func fetch(id int) {
	fetchMu.Lock()
	defer fetchMu.Unlock()

	if !lastFetch.IsZero() {
		wait := time.Until(lastFetch.Add(time.Second))

		if wait > 0 {
			time.Sleep(wait)
		}
	}

	lastFetch = time.Now()

	fmt.Printf(
		"Task %d: Fetch started at %s\n",
		id,
		lastFetch.Format("15:04:05.000"),
	)

	// Simulate an API request.
	time.Sleep(200 * time.Millisecond)

	fmt.Printf("Task %d: Fetch finished\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			fetch(id)
		}(i)
	}

	wg.Wait()

	fmt.Println("All tasks completed")
}
