package main

import (
	"fmt"
)

func main(){
	var ms *myStruct
	// pointer u don't initialise is actually going to be initialised for u, and it goona hold the value nil
	fmt.Println(ms)
	// ms  = &myStruct{foo:42}
	 ms = new(myStruct)


	// get field value
	// (*ms).foo = 20
	//  SAME
	ms.foo = 42
	fmt.Println(ms)

	// defrencing operator actually has a low precedence than dot(.) operator
}

type myStruct struct{
	foo int
}