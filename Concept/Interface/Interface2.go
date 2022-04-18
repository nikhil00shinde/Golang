package main 


import (
	"fmt"
)

func main(){
	 myInt := IntCounter(0) //casting int to IntCounter
	 var inc Incrementer = &myInt
	 for i := 0; i < 10; i++ {
		 fmt.Println(inc.Increment())
	 }
}

// You don't have to use struct to implement interfaces, you can use any type of custom type 
// But any type do have control over that I can create, I can methods to it , If I can add method to it I can implement interfaces with it

type Incrementer interface{
	Increment() int
}
// type alias for an integer
type IntCounter int 

func (ic *IntCounter) Increment() int{
   *ic++
	 return int(*ic)
}