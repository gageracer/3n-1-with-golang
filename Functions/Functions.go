package main

import (
	"fmt"
)

func main() {
	sayMessage("Wow")
	greting := "Hello"
	name := "Stacey"
	sayGreeting(&greting, &name)
	fmt.Println("The sum is", sum(1, 2, 3, 4, 5))

	// functions are types, d is assigned to the return of the func
	d, err := divide(5, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()

}

func sayMessage(msg string) {
	fmt.Println(msg)
}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
}

func sum(values ...int) (result int) {
	fmt.Println(values)
	result = 0
	for _, v := range values {
		result += v
	}
	return
}

// Getting/returning the error in a function
func divide(a, b float64) (float64, error) {
	if b == 0. {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

// struct
type greeter struct {
	greeting string
	name     string
}

// method for the struct, gives the struct a function
// that can access to its variables
func (g *greeter) greet() { // it would make sense to have the greeter a pointer
	fmt.Println(g.greeting, g.name)
}
