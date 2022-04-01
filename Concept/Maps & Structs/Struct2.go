package main

import (
	"fmt"
)

// Go doesnot use OOP concept, so instead of inheritence we have composition

type Animal struct{
	Name string
	Origin string
}

type Bird struct{
	Animal
	SpeedKPH float32
	CanFly bool
}

func main(){
	//  Bird has a animal like characteristics (has-a relationship)
	
	// LITERAL SYNTAX -> Embedding
	b := Bird{
		Animal: Animal{Name:"Emu",Origin: "Australia"},
		 SpeedKPH: 48,
	 CanFly: false,
	}


	//  b.Name = "Emu"
	//  b.Origin = "Australia"
	//  b.SpeedKPH = 48
	//  b.CanFly = false
	 fmt.Println(b)
}