package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)

		time.Sleep(500 * time.Millisecond)

		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	var wg sync.WaitGroup

	// Start two workers.
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs.
	for i := 1; i <= 5; i++ {
		jobs <- i
	}

	close(jobs)

	// Wait for workers to finish, then close results.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results.
	for result := range results {
		fmt.Println("Result:", result)
	}

	fmt.Println("All jobs completed")
}
