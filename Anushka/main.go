package main

import (
	"fmt"
	"golang-starter/myutil"
	stringpractice "golang-starter/string_practice"
)

func main() {
	fmt.Println("Hello World from Golang")
	myutil.Printmessage("Hello dosto , kya haal hai ? ")

	var name string = "Anushka Yadav"
	fmt.Println(name)

	var age int = 23
	fmt.Println(age)

	const phone = 7395058340
	fmt.Println(phone)

	var salary float64 = 54340.00
	fmt.Println(salary)

	company := "Sunware Technologies"
	fmt.Println(company)

	result := stringpractice.ReplaceChar("Rum", 'M')
	fmt.Println(result)

	var userInput string
	fmt.Print("Enter a word: ")
	fmt.Scan(&userInput)
	bytes, chars := stringpractice.CountState(userInput)
	fmt.Printf("For '%s' -> Bytes: %d, Characters: %d\n", userInput, bytes, chars)
}
