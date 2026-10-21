package main

import "fmt"

func main() {
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	j := 1
	for {
		fmt.Println(j)
		j++
		if j > 5 {
			break
		}
	}

	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}
	for _, value := range numbers {
		fmt.Println(value)
	}
	for index := range numbers {
		fmt.Println(index)
	}
}
