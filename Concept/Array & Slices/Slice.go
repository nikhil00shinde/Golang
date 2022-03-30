package main

import (
	"fmt"
)


func main(){
	// slice are reference type
	a := []int{1,2,3}

  b := a
	b[1] = 10
	fmt.Println(a)
	fmt.Println(b)  	
	
	fmt.Printf("Length: %v\n",len(a))
	fmt.Printf("Capacity: %v\n",cap(b))


  // can also work with array
	a := [...]int{1,2,3,4,5,6,7,8,9,10}
	a := []int{1,2,3,4,5,6,7,8,9,10}

	b := a[:]   // slice of all elements
	c := a[3:]  // slice from 4th element to end
	d := a[:6]  // slice first 6 elements
	e := a[3:6] // slice the 4th, 5th and 6th elements

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)


	// to make slice(type,length,capacity)
	a := make([]int,3,100)

	fmt.Println(a)
	
	fmt.Printf("Length: %v\n",len(a))
	fmt.Printf("Capacity: %v\n",cap(a))
}