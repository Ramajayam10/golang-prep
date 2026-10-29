package main

import (
	"fmt"
	"time"
)

func worker(ch chan int) {
	fmt.Println("Worker waiting...")
	value := <-ch
	fmt.Println("Received:", value)
}

func main() {
	ch := make(chan int)

	go worker(ch)

	time.Sleep(1 * time.Second)

	fmt.Println("Main finished")
}
