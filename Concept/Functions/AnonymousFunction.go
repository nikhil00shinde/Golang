package main

import (
	"fmt"
)

func main(){
	// Anonymous function
	// func() {
	// 	fmt.Println("Hello Go!")
	// }() // immediately invoked function

	// for i:=0; i < 5; i++ {
	// 	func (i int)  {
	// 	      fmt.Println(i)	
	// 	}(i)
	// }

	// variable function
	f := func(){
		fmt.Println("Hello Go!")
	}
	f()
}