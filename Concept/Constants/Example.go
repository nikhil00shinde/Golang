package main

import "fmt"

const (
	isAdmin = 1 << iota
	isHeadquater
	canSeeFiniancial
	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main(){

	var roles byte = isAdmin | canSeeFiniancial | canSeeEurope
	fmt.Printf("%b\n",roles)
	fmt.Printf("isAdmin? %v\n",isAdmin & roles == isAdmin)
	fmt.Printf("isHQ? %v\n",isHeadquater & roles == isHeadquater)

}