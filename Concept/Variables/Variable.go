// Variable declaration
// Redeclaration and shadowing(prcedence of variable in scope is called shadowing)
// visibility
// Naming convention
// Type conversion

package main

import "fmt"

var (
i int = 43
name string = "Nikhil"
d float32 = 56.5
)

func main(){
//  var i int 
//   i = 42

//  var i int = 42

//    i := 42 

// var i int
// i = 42

// var j int = 23
// var j float32 = 23

//  k := 2

	fmt.Printf("%v, %T",d,d) //print in string formatter
}



// -----------

//  - If local variable is declared and not used then it will throw an error(compiletime error)