// arr, slice, map

package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	slice := []int{10, 20, 30}

	fmt.Println(arr)
	fmt.Println(slice)

	slice = append(slice, 40)
	slice = append(slice, 50)

	fmt.Println(slice)
	fmt.Println(len(slice))

	for _, value := range slice {
		fmt.Println(value)
	}
	numbers := []int{10, 20, 30, 40, 50, 60, 70}
	fmt.Println(numbers[0])
	fmt.Println(numbers[6])
	fmt.Println(numbers[1:4])
	fmt.Println(numbers[4:])

	player := map[string]string{
		"name":    "Sachin",
		"country": "India",
		"age":     "53",
	}
	player["city"] = "Mumbai"
	player["name"] = "Sachin Tendulkar"
	city, ok := player["city"]
	fmt.Println(ok, city, player["name"])

	runs := make(map[string]int)

	runs["sachin"] = 200
	runs["india"] = 400
	runs["Guptil"] = 0
	kohli, okKohli := runs["kohli"]
	guptil, ok := runs["Guptil"]
	fmt.Println(runs["sachin"], guptil, kohli, okKohli, ok)

	runs["kohli"] = 183
	runs["india"] = 401
	delete(runs, "Guptil")
	delete(runs, "kohlii")
	fmt.Println(runs)

	for player, score := range runs {
		fmt.Println(player, score)
	}
}
