package main

import (
	"fmt"
)

func main(){
	
	// f := func(){
	// 	fmt.Println("Hello Go!")
	// }
	// f()
  
	// function signature
	var f func() = func(){
		fmt.Println("Hello Go!")
	}
	f()
	 
 //example
 var divide func(float64,float64) (float64,error)

 divide = func(a, b float64) (float64, error) {
          if b == 0.0 {
						return 0.0 , fmt.Errorf("Cannot divide by zero")
					} else{
						return a / b , nil
					}
 }

 d,err := divide(5.0,3.0)
 if err != nil{
	 fmt.Println(err)
	 return
 }
 fmt.Println(d)
}