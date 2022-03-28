package main

import "fmt"

func main()  {

	// var n bool = true
	// Uninitialized variable will have 0(zero) value
	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v %T\n",n,n)
	fmt.Printf("%v %T\n",m,m)

}