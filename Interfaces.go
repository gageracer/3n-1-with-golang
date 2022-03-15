package main

import (
	"fmt"
)

func main() {
	// Writer
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Helllo"))

	// Incrementer
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i:=0;i<10;i++{
		fmt.Println(inc.Increment())
	}


}

// Interface
type Writer interface {
	Write([]byte) (int,error)
}
// Struct
type ConsoleWriter struct {}
// Method for the Struct
func (cw ConsoleWriter) Write(data []byte) (int,error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// Interface for The Increment
type Incrementer interface {
	Increment() int
}
// a new type with name IntCounter
type IntCounter int
// Method
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
 
 