package main //entry point

import (
	"fmt"
)

func main() {
	   s := sum("The sum of ",1,2,3,4,5)
		 fmt.Println(s)
}

func sum(msg string,values ...int)  (result int){
       fmt.Println(msg,values)
			 for _,v := range values{
				 result += v
			 }

			 return
}