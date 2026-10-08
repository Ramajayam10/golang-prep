package main
import "fmt"

type Car struct {
	brand string
	model string
}

type Vehicle interface {
	start(string)
	stop()
}

func (car Car) start(c string) {
	fmt.Printf("Car is starting, %s", c)
}

func (car Car) stop() {
	fmt.Println("Car is stopping")
}

func main() {
	car1 := Car{brand: "Toyota", model: "Fortuner"}
	car1.start(car1.brand)
}