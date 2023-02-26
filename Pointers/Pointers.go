package main

import (
	"fmt"
)

func main() {
	a:=42
	b := &a
	fmt.Println(a,b,&a)
	a = 27
	fmt.Println(a,*b)
	// when assign primitives to another variable, 
	// it creates a new copy
	// to get their address, you need to use pointers
	
	// for maps and slices, assigning to another variable
	// gives the address of the map/slice because both are
	// pointers already
}