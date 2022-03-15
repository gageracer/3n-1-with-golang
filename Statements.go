package main

import (
	"fmt"
)

func main() {
	number := 50
	guess := 11
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it")
		}

	}

	switch 2{
	case 1,3,4:
		fmt.Println("one")
	case 2:
		fmt.Println("Right in two")
	default:
		fmt.Println("not one or two")
	}
	tagy := 3
	switch {
	case tagy < 10:
		fmt.Println("less than 10")
		fallthrough // continues the switch case and runs the next case anyway
	case tagy > 10:
		fmt.Println("More than 10")
	default:
		fmt.Println("it is 10")
	}

	var inte interface{} = true
	switch inte.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("another type")
		
	}
}