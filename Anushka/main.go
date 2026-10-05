package main

import (
	"fmt"
	"golang-starter/myutil"
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
}
