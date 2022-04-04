package main

import (
	"fmt"
)

func main(){
	// a,b:=1,2
	// ans := a /  b
	// fmt.Println(ans)

	// fmt.Println("start")
	// panic("something bad happened")
	// fmt.Println("end")
 

	// panic happened after defer statements are executed
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something bad happened")
	fmt.Println("end")
}