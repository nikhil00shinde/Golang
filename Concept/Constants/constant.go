package main

import (
	"fmt"
)

func main()  {
	// const myConst int = 42  //should be assign at compile time not run time

	// IMPLICIT CONVERSION
   const a = 42
	 var b int16 = 27
	fmt.Printf("%v, %T",a+b,a+b)
}