package main 

import (
	"fmt"
)

func main() {
	g := greeter{
		greeting: "Hello",
		name:"Go",
	}
	//method invoking
	g.greet()
	fmt.Println("The new name is: ",g.name)
	g.greetP()
	fmt.Println("The new name is: ",g.name)
 
}
type greeter struct{
	greeting string
	name string
}

// method is function that is executing in an known context(any type)
func (g greeter) greet(){
	fmt.Println(g.greeting,g.name)
	g.name  = ""
}

func (g *greeter) greetP(){
	fmt.Println(g.greeting,g.name)
	g.name  = ""
}
 