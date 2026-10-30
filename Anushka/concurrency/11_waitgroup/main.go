package main

import (
	"fmt"
	"sync"
	"time"
)

// One single function reused by all goroutines!
func sendEmail(userName string, wg *sync.WaitGroup) {
	defer wg.Done() // Step 3: When this function finishes, decrement counter (-1)

	fmt.Printf("Sending welcome email to %s...\n", userName)
	time.Sleep(1 * time.Second) // Simulating network email sending
	fmt.Printf(" Email successfully sent to %s!\n", userName)
}

func main() {
	var wg sync.WaitGroup

	users := []string{"Anushka", "Rahul", "Priya"}

	fmt.Println("Starting email blast...")

	for _, name := range users {
		wg.Add(1)               // Step 1: Add +1 to counter for each person
		go sendEmail(name, &wg) // Step 2: Launch in background
	}

	// Step 4: Freeze main() right here until counter reaches 0!
	wg.Wait()

	fmt.Println(" All welcome emails have been sent! Process finished.")
}
