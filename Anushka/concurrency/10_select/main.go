package main

import (
	"fmt"
	"time"
)

func main() {
	uberCh := make(chan string)
	olaCh := make(chan string)

	// Uber driver is searching (takes 3 seconds)
	go func() {
		time.Sleep(3 * time.Second)
		uberCh <- " Uber driver accepted!"
	}()

	// Ola driver is searching (takes 1 second)
	go func() {
		time.Sleep(5 * time.Second)
		olaCh <- " Ola driver accepted!"
	}()

	// SELECT: Listens to ALL cases at the exact same time!
	select {
	case cab := <-uberCh:
		fmt.Println("RIDE CONFIRMED:", cab)

	case cab := <-olaCh:
		fmt.Println("RIDE CONFIRMED:", cab)

	case <-time.After(2 * time.Second):
		fmt.Println(" TIMEOUT: No drivers found within 2 seconds. Cancelling request!")
	}
}
