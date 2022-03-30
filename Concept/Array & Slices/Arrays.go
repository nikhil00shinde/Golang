package main

import (
	"fmt"
)


func main(){
  
	// grades := [3]int{1,2,3}
	// grades := [...]int{1,2,3}

	var students [3]string
	fmt.Printf("Students: %v\n",students)

	students[0] = "Pain"
	students[1] = "Rinne"
  students[2] = "Japan"

	fmt.Printf("Students: %v\n",students)

	fmt.Printf("Number of Students: %v\n",len(students))




}