package main

import "fmt"




func main() {
	// numbers := []int{1, 2, 3, 4, 5}
	// 	numbers := make([]int,3,5)
	// 	numbers = append(numbers,6,8,9,12,13,141,500)
	// 	fmt.Println("The numbers are ", numbers)
	// 	fmt.Printf("The data type of numbers is %T\n", numbers)
	// 	fmt.Println("Length is :", len(numbers))
	// 	fmt.Println("The capacity is : ", cap(numbers))
	numbers := []int{10, 20, 30}
	numbers = append(numbers, 40, 50, 60)
	fmt.Println("the numbers are :", numbers)
	fmt.Println("The length of numbers is :", len(numbers))
	i := 0
	for i < len(numbers) {
		fmt.Println(numbers[i])
		i = i + 1
	}
	for index, value := range numbers {
		fmt.Println(index, value)
	}

	sum := 0
	i = 0
	for i < len(numbers){
		sum = sum + numbers[i]
		fmt.Println(sum)
		i++
	}
	fmt.Println("the total sum of the number is : ",sum)

	largest := numbers[0]
	i = 0
	for i < len(numbers){
		if numbers[i]> largest{
			largest = numbers[i]

		}
		i++
	}
	fmt.Println("the largest number is ",largest)


	smallest := numbers[0]
	i = 0

	for i < len(numbers){
		if numbers[i] < smallest{
			smallest = numbers[i]
		}
		i++
	}
	fmt.Println("The smallest is : ",smallest)
}

// Whenever we are creating slice like ( numbers := []int) it will through error bcz it need a initial value

//  * so if we want to create any slice without providing initial value then we have to use make function
//  *make() function is used to create a slice with a specific length and capacity
