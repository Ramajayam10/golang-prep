package main

import "fmt"

func main() {
	ch := make(chan int)

	fmt.Println("Before sending")

	ch <- 10

	fmt.Println("After sending")
}
