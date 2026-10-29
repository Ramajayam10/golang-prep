package main

import (
	"fmt"
	"sync"
	"time"
)

func task1(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Task 1 started")
	time.Sleep(2 * time.Second)
	fmt.Println("Task 1 finished")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Task 2 started")
	fmt.Println("Task 2 finished")
}

func task3(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Task 3 started")
	fmt.Println("Task 3 finished")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go task1(&wg)
	go task2(&wg)
	go task3(&wg)

	wg.Wait()

	fmt.Println("All tasks done!")
}
