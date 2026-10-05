package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Please enter your name :")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Welcome ", name)

	fmt.Println("Pls enter your name :")
	var name string
	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n')
	fmt.Println("Hello! ", name)

	fmt.Println("Enter your company name: ")
	var company string
	reader = bufio.NewReader(os.Stdin)
	company, _ = reader.ReadString('\n')
	fmt.Println("Ohk So you work in :", company)
}
