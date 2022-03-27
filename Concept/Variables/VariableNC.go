package main

import "fmt"

// var i int = 42 //lower case variable are scoped to package
var I int = 42 //exported from the package n globally visible


func main(){
	//block scope
	var i int = 43

	j := 12

	fmt.Println(i,j)
}


//Can't redeclare variables, but can shadow them
//All variables must ne used

// VISIBILTY
   //lower case first letter for package scopr
	 //upper case first letter to export
	 //no private scope

// Naming convention
  //Pascal or camelCase
	   // Capitalize acronyms(HTTP,URL)


// TYPE CONVERSION
// destinationType(variable)
// use strconv package for strings