package main

import (
	"fmt"
)

func main(){
	a := []int{}
	fmt.Println(a)
	fmt.Printf("Length: %v\n",len(a))
	fmt.Printf("Capacity: %v\n",cap(a))

	a = append(a, 1)
	fmt.Println(a)
	fmt.Printf("Length: %v\n",len(a))
	fmt.Printf("Capacity: %v\n",cap(a))

	a = append(a, 2,3,4,5)
	// like javascript we can spread
	a = append(a, []int{6,7,8,9}...)
	fmt.Println(a)
	fmt.Printf("Length: %v\n",len(a))
	fmt.Printf("Capacity: %v\n",cap(a))
}