package main

import (
	"fmt"
	"time"
)

func task1() {
	fmt.Println("Task 1 started")
	time.Sleep(2 * time.Second)
	fmt.Println("Task 1 finished")
}

func task2() {
	fmt.Println("Task 2 started")
	fmt.Println("Task 2 finished")
}

func main() {
	go task1()
	go task2()

	time.Sleep(3 * time.Second)
	fmt.Println("Main finished")
}
