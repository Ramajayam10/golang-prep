package main

import "fmt"

func main() {
    var age int = 25
    var price float64 = 99.99
    var name string = "Rick Grimes"
    var isStudent bool = true

	isHosteller := true
	isHosteller = false

	fmt.Println(isHosteller)
    fmt.Println(age)
    fmt.Println(price)
    fmt.Println(name)
    fmt.Println(isStudent)
	if age >= 18 {
		fmt.Println("Adult")
	} else if age < 13 {
		fmt.Println("Child")
	} else {
		fmt.Println("Minor")
	}

	// var (
	// 	name string = "Ajith"
	// 	age  int    = 54
	// 	state string = "Tamilnadu"
	// )
	// x, y := 10, 20
}
