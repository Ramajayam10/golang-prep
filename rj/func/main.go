package main

import (
	"encoding/json"
	"fmt"
)

func multiply(a, b int) (result int) {
	result = a * b
	return result
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func getUser() (string, error) {
	return "ram", nil
}

type User struct {
	Name string `json:"name"`
	Age  int	`json:"age"`
	City string	`json:"city"`
}

func main() {
	result := multiply(1, 1)
	fmt.Println(result)
	name, geterr := getUser()

	fmt.Println(name)
	fmt.Println(geterr)
	res, err := divide(10, 10)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

	u := User{
		Name: "Sivasaami",
		Age: 45,
		City: "Thoothukudi",
	}
	//struct to json
	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	//json to struct
	jsonData := `{"name":"Sivasaami","age":45,"city":"Thoothukudi"}`
	var user User
	jsonerror := json.Unmarshal([]byte(jsonData), &user)

	if jsonerror!=nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user.Name)
	fmt.Println(user.Age)
}
