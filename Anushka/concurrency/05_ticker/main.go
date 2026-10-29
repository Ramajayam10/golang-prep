package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	countdown := 5
	fmt.Println("Mission Control: Preparing for launch...")

	for range ticker.C {
		fmt.Printf("%d seconds remaining...\n", countdown)
		countdown--

		if countdown == 0 {
			fmt.Println("BLAST OFF! The rocket has been launched into space!")
			break
		}
	}
}
