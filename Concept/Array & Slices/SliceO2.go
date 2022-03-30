package main

import (
	"fmt"
)

func main(){
	a := []int{1,2,3,4,5}
	// remove element from starting
	// b := a[1:]

	// remove element from ending
	// b := a[:len(a)-1]


	// remove element from middle
	b := append(a[:2],a[3:]...)
	fmt.Println(b)
}