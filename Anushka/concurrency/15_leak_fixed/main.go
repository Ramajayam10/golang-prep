package main

import "fmt"

func worker(done chan struct{}) {
	fmt.Println("Worker started")

	<-done

	fmt.Println("Worker stopped")
}

func main() {
	done := make(chan struct{})

	go worker(done)

	close(done)

	fmt.Println("Main finished")
}
