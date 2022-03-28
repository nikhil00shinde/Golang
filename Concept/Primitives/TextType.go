package main

import (
	"fmt"
)

func main()  {
	 
	  s := "this is a string"
		// s2 := "this is also a string"

		b := []byte(s)
		fmt.Printf("%v, %T\n",b,b)


		// Rune -> are type alias for int32
		var r rune = 'a'
		// r := 'a'
		fmt.Printf("%v, %T\n",r,r)
}


// Strings are immutable 