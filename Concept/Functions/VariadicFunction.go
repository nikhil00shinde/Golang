package main //entry point

import (
	"fmt"
)

func main() {
	   s := sum("The sum of ",1,2,3,4,5)
		 fmt.Println(*s)
}


// func sum(msg string,values ...int) int{
//        fmt.Println(msg,values)
// 			 result := 0 
// 			 for _,v := range values{
// 				 result += v
// 			 }
// 			 return result
// }

func sum(msg string,values ...int) *int{
       fmt.Println(msg,values)
			 result := 0 
			 for _,v := range values{
				 result += v
			 }
			//  returning a value that is generated on local stack, it will automatically promote variable for you to be on a shared memory in the computer,it called heap memory
			 return &result
}