package main

import (
	"fmt"
	"time"
)

// func main() {
// 	fmt.Println("Timer started")

// 	timer := time.NewTimer(3 * time.Second)

// 	fmt.Println("Waiting for timer...")

// 	alarmTime := <-timer.C
// 	fmt.Println("The alarm rang at:", alarmTime)
// }


func main() {
	fmt.Println("🍕 Order placed! You have 3 seconds to cancel.")
	fmt.Println("👉 Press [ENTER] to CANCEL the order:")

	timer := time.NewTimer(3 * time.Second)
	userCancelled := make(chan bool)

	//Listen for user pressing Enter:
	go func() {
		var input string
		fmt.Scanln(&input)
		userCancelled <- true
	}()

	// The Race: Whichever happens first wins!
	select {
	case <-userCancelled:
		// User pressed Enter before 3 seconds!
		timer.Stop()
		fmt.Println("✅ Order Cancelled! Full refund issued.")

	case <-timer.C:
		// 3 seconds passed 
		fmt.Println("\n⏰ 3 seconds expired! Cancellation window closed.")
		fmt.Println("👨‍🍳 Chef started cooking your pizza!")
	}
}
