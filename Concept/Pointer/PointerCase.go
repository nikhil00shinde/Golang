package main

import (
	"fmt"
)

func main(){
	// Primitive datatype struct and array uses copy not pointer
	// a := [3]int{1,2,3}
	// b is a copy of array
	// b := a
	// fmt.Println(a,b)
	// a[1] = 42
	// fmt.Println(a,b)

	// SLICE
	// a := [3]int{1,2,3}
	// // slice is a projection of an underline array 
	// b := a 
	// fmt.Println(a,b)
	// a[1] = 42
	// fmt.Println(a,b)

	// MAP
	a := map[string]string{
		"foo":"bar",
		"baz":"bar",
	}
	b := a
	fmt.Println(a,b)
	a["foo"] = "qux"
	fmt.Println(a,b)

	// slices and map underline both uses pointer to data
}