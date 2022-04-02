package main

import (
	"fmt"

)

func main(){

  // initializer; 
	switch i := 2 + 3;i {
	case 1, 5, 10:
		      fmt.Println("One, Five, Ten") 
	case 2, 4, 6:
		      fmt.Println("Two, Four, Six") 
	default:
		      fmt.Println("another number")
	}

	// tagless syntax
	// i := 10
	// switch {
	// case i <= 10:
	// 	   fmt.Println("less than or equal to ten")
	// 		 fallthrough
  // case i <= 20:
	// 	   fmt.Println("less than or equal to twenty")
	// default:
	// 	  fmt.Println("greater than twenty")
	// }

	// type switch
	var i interface{} = [3]int{}
	switch i.(type){
	case int:
		fmt.Println("i is an int")
		break
		fmt.Println("This will be print too")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is string")
	case [3]int:
		fmt.Println("i is [3]int")
	default:
		fmt.Println("i is another type")
	}
}