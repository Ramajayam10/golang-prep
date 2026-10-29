package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			mu.Lock()
			count++
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			mu.Lock()
			count++
			mu.Unlock()
		}
	}()

	wg.Wait()

	fmt.Println("Count:", count)
}
