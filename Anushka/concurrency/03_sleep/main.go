package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Goroutine started")
		time.Sleep(3 * time.Second)
		fmt.Println("Goroutine finished")
	}()

	fmt.Println("Main finished")
}
