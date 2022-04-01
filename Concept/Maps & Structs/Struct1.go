package main

import (
	"fmt"
)

func main(){
	// not reference pass -> passing copies only
	aDoctor := struct{name string}{name:"Joe Pertwee"}
	// anotherDoctor := aDoctor

	// pass by reference
	anotherDoctor := &aDoctor


  anotherDoctor.name = "Tom Baker"

	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
}