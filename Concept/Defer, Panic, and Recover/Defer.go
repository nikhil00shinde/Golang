package main 

import (
	"fmt"
)

func main(){
	defer fmt.Println("first")
	defer fmt.Println("middle")
	defer fmt.Println("last")
	// defer keyword -> it actually execute any function passed to it, after function finishes it's final statement but before it actually returns
	// defer function works in LIFO order (last in first out)

	a := "start"
	defer fmt.Println(a) //start
	// when u defer a function like this it actually argument at the time the defer is called not the time the called function executed
	a = "end"
}