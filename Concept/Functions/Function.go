package main //entry point

import (
	"fmt"
)

func main() {
   fmt.Println("Hello, playground")
	 sayMessage1("Hello Go!\n\n")
	 for i := 0; i < 5; i++ {
		 sayMessage2("Hello Go!",i)
	 }
}
 
func sayMessage1(msg string) {
	fmt.Println(msg)
}

func sayMessage2(msg string,idx int){
  fmt.Println(msg)
	fmt.Println("The value of the index is",idx)
}

func sayGreeting(greeting,name string){
     fmt.Println(greeting,name )
}