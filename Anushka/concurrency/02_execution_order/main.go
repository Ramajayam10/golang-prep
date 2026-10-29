package main

import (
	"fmt"
	"time"
)

func task(name string) {
	fmt.Println(name, "started")
	time.Sleep(1 * time.Second)
	fmt.Println(name, "finished")
}

func main() {
	go task("Task 1")
	go task("Task 2")
	go task("Task 3")

	time.Sleep(2 * time.Second)
	fmt.Println("Main finished")
}
