package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a, b int) (result int) {
	result = a * b
	return
}

func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Denominator can't be zero")
	}
	result := a / b
	return result, nil
}

func square(a int) int {
	return a * a
}

func cube(a int) int {
	return a * a * a
}

func Iseven(a int) bool {
	return a%2 == 0
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func isAdult(age int) bool {
	return age >= 18
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func greet(name string) string {
	return "Hello! " + name + " Welcome to the Golang Learning Journey"
}

func remainder(a, b int) (float64, float64) {
	quoteint := a / b
	remainder := a % b
	return float64(quoteint), float64(remainder)
}

func sumToNumber(n int) int {
	sum := 0
	i := 1
	for i <= n {
		sum += i
		i += 1
	}
	return sum
}

func factorial(a int) int {
	factor := 1
	i := 1
	for i <= a {
		factor *= i
		i += 1
	}
	return factor
}

func main() {
	add_result := add(5, 3)
	fmt.Println("Addition of two number is :", add_result)

	sub_result := subtract(9, 5)
	fmt.Println("The subtraction of two numbers is :", sub_result)

	multiply_result := multiply(3, 4)
	fmt.Println("The multiplication of two numbers is :", multiply_result)

	division_result, err := division(13, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The division of two number is :", division_result)
	}

	// Square
	fmt.Println("Please enter the number")
	var num_square int
	fmt.Scan(&num_square)
	square_result := square(num_square)
	fmt.Println("The square of provided number is :", square_result)

	// cube
	fmt.Println("Enter your number pls :")
	var num_cube int
	fmt.Scan(&num_cube)
	cube_result := cube(num_cube)
	fmt.Println("The cube of your number is :", cube_result)

	// check even
	fmt.Printf("Please enter your number to check : ")
	var IsEven_num int
	fmt.Scan(&IsEven_num)
	check_even := Iseven(IsEven_num)
	if check_even {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	// Check max
	fmt.Println("PLease enter your 1st number")
	var num_first int
	fmt.Scan(&num_first)

	fmt.Println("Please enter your 2nd number")
	var num_second int
	fmt.Scan(&num_second)

	max_result := max(num_first, num_second)
	fmt.Println("The max number is ", max_result)

	// Check min
	fmt.Println("Please enter your 1st number")
	var first_number int
	fmt.Scan(&first_number)

	fmt.Println("Please enter your 2nd number")
	var second_number int
	fmt.Scan(&second_number)

	min_result := min(first_number, second_number)
	fmt.Println("The min number is ", min_result)

	// wheter adult or not
	fmt.Println("Please Enter your Age :")
	var age int
	fmt.Scan(&age)

	check_adult := isAdult(age)
	fmt.Println("Are you an adult ?", check_adult)

	// Temperature
	fmt.Println("Please enter your temp in celcius :")
	var temp float64
	fmt.Scan(&temp)

	frrTemp := celsiusToFahrenheit(temp)
	fmt.Println("The temperature in fahrenheit is ", frrTemp)

	// Greet
	fmt.Println("Please provide your good name :")
	var name string
	fmt.Scan(&name)
	greet_result := greet(name)
	fmt.Println(greet_result)

	// quoteient and remainder
	fmt.Println("Enter your number :")
	var num1 int
	fmt.Scan(&num1)

	fmt.Println("Enter your 2nd num :")
	var num2 int
	fmt.Scan(&num2)

	quotient, remainder := remainder(num1, num2)
	fmt.Println("The quotient is ", quotient)
	fmt.Println("The remainder is ", remainder)

	fmt.Println("Enter your number to check sum :")
	var num_sum int
	fmt.Scan(&num_sum)
	sum_to_number := sumToNumber(num_sum)
	fmt.Println("The sum of the number is ", sum_to_number)

	fmt.Println("Enter your number :")
	var numFactorial int
	fmt.Scan(&numFactorial)
	factorialRes := factorial(numFactorial)
	fmt.Println("The factorial of the number is ", factorialRes)
}
