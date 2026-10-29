package main

import "fmt"

func calculate(ch chan int) {
	result := 10 + 20
	ch <- result
}

func main() {
	ch := make(chan int)

	go calculate(ch)

	result := <-ch

	fmt.Println("Result:", result)
}
