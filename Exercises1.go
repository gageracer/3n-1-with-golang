package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true
var yy int

type sik int

func main() {
	var x sik
	s := fmt.Sprintf("%v %T,\t%v %T,\t%v %T\n", x, x, y, y, z, z)
	fmt.Print(s)
	fmt.Printf("%v %T\n", x, x)
	x = 42
	fmt.Printf("%v %T\n", x, x)
	yy = int(x)
	fmt.Printf("%v %T",yy,yy)
}
