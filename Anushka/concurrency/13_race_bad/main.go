package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			count++
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			count++
		}
	}()

	wg.Wait()

	fmt.Println("Count:", count)
}
