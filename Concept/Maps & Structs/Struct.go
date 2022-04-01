package main

import (
	"fmt"
)

type Doctor struct{
	number int
	actorName string
	companions []string
}

func main()  {
	// Field name
	// Initilizer
	aDoctor := Doctor{
		number:3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
  // VALID SYNTAX (Positional syntax)
	// aDoctor := Doctor{
	// 	3,
	// 	"Jon Pertwee",
	// 	[]string{
	// 		"Liz Shaw",
	// 		"Jo Grant",
	// 		"Sarah Jane Smith",
	// 	},
	// }

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.companions[1])

	// Anonymous Struct
  bDoctor := struct{name string}{name:"Rinne"}
  fmt.Println(bDoctor)
}