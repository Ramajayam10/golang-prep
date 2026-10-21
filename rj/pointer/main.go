package main

import "fmt"

func change(x *int) {
	*x = *x + 10
}

type User struct {
	name string
	age  int
	city string
}

func main() {
	x := 10
	p := &x
	x = 11
	*p = 20
	fmt.Println(x)

	num := 20
	change(&num)

	fmt.Println(num)

	user := User{
		name: "George David",
		age:  25,
		city: "Aluva",
	}
	u := &user

	u.age = 26
	bday(&user)
	fmt.Println(u.age)
	user.birthday()
	fmt.Println(u.age)
}

func bday(user *User) {
	user.age++
}

func (user *User) birthday() {
	user.age++;
}