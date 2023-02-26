package main

import (
	"fmt"
	"log"
)

func main() {
	// defers run as last in first out
	fmt.Println("Start")
	defer fmt.Println("defer 1") // defer makes the function run after it returns
	defer fmt.Println("defer 2")

	// Panic
	panicker()
	// a,b := 1,0
	// ans := a / b // This will give a panic
	// fmt.Println(ans)
	defer fmt.Println("defer 3")
	// panic("something bad happened")
	fmt.Println("end")
}

// with a scoped function panicking and recovering, the upper function is not affected by the error
func panicker() {
	fmt.Println("I'm about to panic!")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			//panic(err) // put this if you want the main function to panic
		}
	}()
	panic("something bad happened!!!\n")
	fmt.Println("done panicking")
}
