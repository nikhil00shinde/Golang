package main

import (
	"fmt"
)

func main(){
	// a := 42
	 // pass by value not address
	// b := a
	// fmt.Println(a,b)
	// a = 27
	// fmt.Println(a,b)

  // CREATING POINTER
	var a int = 42
	var b *int = &a
	fmt.Println(a,*b)
  a = 29
	*b = 53
	fmt.Println(a,*b)
}