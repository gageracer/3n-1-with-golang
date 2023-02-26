package main

import (
	"fmt"
)

func main() {
	for i,j:= 0,1; i<5 ; i,j = i+1,j+1 {
		fmt.Println(i,j)
	}
	i:=0
	// while loop
	for i<5 {
		fmt.Println(i)
		i++
	}
	// Infinite while loop
	for {
		fmt.Println("INFINITE")
		if i >4 {
			break // put a special break scenario, can also use continue but will loop forever
		} 
	}

	// Label to return or break to
	Loop:
		for {
			for{
				fmt.Println("Double Nested Infinite For Loop?")
				break Loop
			}
		}
	
	// For range loop
	s := []int{1,2,3,4,5}
	for k,v := range s{ // you can _,v to skip k usage with compiler
		fmt.Println(k,v) // key and value, can be used for maps,strings as well
	}


}