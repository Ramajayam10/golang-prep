package main

import "fmt"

func multiply(a, b int) (result int) {
	result = a * b
	return result
}

func getUser() (string, int) {
	return "ram", 26
}
func main () {
	result := multiply(1, 1);
	fmt.Println(result)
    name, age := getUser()

    fmt.Println(name)
    fmt.Println(age)
}