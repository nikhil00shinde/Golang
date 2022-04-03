package main

import (
	"fmt"
)

func main(){
  // label
  // s := []int{1,2,3}
	// fmt.Println(s,"\n")

	// for range loop 
	// for k,v := range s {
	// 	  fmt.Println(k,v)
	// }

	// map
	// statePopulations := map[string]int{
	// 	"q":123,
	// 	"md":1323,
	// 	"medfe":13313,
	// }
	// for k,v := range statePopulations {
	// 	fmt.Println(k,v)
	// }

	// string
	// s := "hello Go!!!"
	// for k,v := range s{
	// 	fmt.Println(k,string(v))
	// }
   
	//for key
	// s := "hello Go!!!"
	// for k:= range s{
	// 	fmt.Println(k)
	// }

	//for value
	s := "hello Go!!!"
	for _,v:= range s{
		fmt.Println(string(v))
	}
}