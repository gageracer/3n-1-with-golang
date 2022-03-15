package main

import (
	"fmt"
	"reflect"
)
// Enumerated Constant
const (
	_ = iota // 0
	interface2
	interface3
)
// Struct
type Person struct{
	Name string `required max:"100"` //tags
}
type Doctor struct {
	number int
	Person
	companions []string
}

var i string = "31"

func main() {
	a := 1
	b := 1
	c := 8
	// Bitwise Operations
	fmt.Println(a & b)  // 1
	fmt.Println(a | b)  // 1
	fmt.Println(a ^ b)  // 0
	fmt.Println(a &^ b) // 0
	fmt.Println(c >> 1) // 4
	fmt.Println(c << 1) // 16

	// Arrays
	grades := [...]int{97, 85, 73}
	fmt.Printf("Grades: %v\nLength: %v\n", grades, len(grades))
	mygrades := grades // copies a new array
	fmt.Printf("myGrades: %v\nmyLength: %v\n", mygrades, len(mygrades))

	// Slice: Arrays with Capacity. these are all pointing to original
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Length: %v\nCapacity: %v\n", len(slice), cap(slice))
	as := slice[:]   // full copy of it
	bs := slice[3:]  // from index 3 to end copy
	cs := slice[3:6] // from index 3 to 6 copy
	fmt.Printf("First Slice:%v\nSecond Slice:%v\nThird Slice:%v\n", as, bs, cs)
	
	ds := make([]int, 3, 100) // makes a slice with 3 elements and 100 capacity
	fmt.Println(ds, len(ds), cap(ds))
	ds = append(ds, []int{1, 2, 3, 4, 5}...) // spreading works like this
	fmt.Println(ds, len(ds), cap(ds))
	ds = ds[3:] // unshifting
	fmt.Println(ds, len(ds), cap(ds))
	ds = ds[:len(ds)-1] // popping
	fmt.Println(ds, len(ds), cap(ds))
	ds = append(ds[:1],ds[2:]...) // removing middle element
	fmt.Println(ds, len(ds), cap(ds))

	 // Maps: Objects for go with strict types for keys and values
	statePopulations := make(map[string]int,10)
	statePopulations = map[string] int{
		"California" : 343423323,
		"Texas": 3423414214,
		"Florida": 1321314234,
	}
	fmt.Println(statePopulations["Texas"])
	statePopulations["Georgia"] = 31
	delete(statePopulations, "Texas")
	_, ok := statePopulations["Texas"] // ok shows if the key is in the map or not
	fmt.Println(ok)
	fmt.Println(statePopulations)


	// Struct
	aDoctor := Doctor{
		Person: Person{Name: "John Pertwee"},// embedding another struct
		number: 3,
		companions: []string{
			"Ahmet","Ayse","Mehmet",
		},
	}
	// getting the tag
	t := reflect.TypeOf(Person{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

	// gets the name from Person struct
	//aDoctor.name = "Jon Pertwee"
	fmt.Println(aDoctor)
	
	// Anonymus Struct that is scoped
	someDoctor:= struct{name string}{name: "Mahmit"}
	fmt.Println(someDoctor)

}
