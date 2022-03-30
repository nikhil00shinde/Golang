package main

import (
	"fmt"
)


func main(){
  
  // var identityMatrix [3][3]int = [3][3]int{ [3]int{1,0,0} ,[3]int{0,1,0},[3]int{0,0,1}}
  
	// var identityMatrix [3][3]int
	// identityMatrix[0] = [3]int{1,0,0}
	// identityMatrix[1] = [3]int{0,1,0}
	// identityMatrix[2] = [3]int{0,0,1}

	// fmt.Println(identityMatrix)

	a := [...]int{1,2,3}

	//on assignment the new array create
  // b := a

	b := &a
	b[1] = 10
	fmt.Println(a)
	fmt.Println(b)



 
	

  	  
}