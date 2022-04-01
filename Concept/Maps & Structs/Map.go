package main

import (
	"fmt"
)

func main(){
	// built-in function to create map 
  statePopulations := make(map[string]int,10)


	statePopulations = map[string]int{
		"California": 39250017,
		"Texas": 27862596,
		"Florida":20612439,
		"New York":19745289,
		"Pennsylvania":12802503,
		"Illinois":12801539,
		"Ohio":11614373,
	}
	// key has to be tested for equality
	// Slices, maps cannot be tested for equality
	// m := map[[]int]string{} slice cannot be a key to a map

	m := map[[3]int]string{} // array is valid key type
	fmt.Println(statePopulations)
	//get
	fmt.Println(statePopulations["Ohio"])

	//add
	statePopulations["Georgia"] = 10310371
  
	//delete
	delete(statePopulations,"Georgia")
	fmt.Println(m)

  // to check if key is present (comma ok method)
  _,ok := statePopulations["Oho"]
   fmt.Println(ok)
  
	//  length
	fmt.Println(len(statePopulations))
  
	// passed by reference
	sp := statePopulations
  delete(sp,"Ohio")
}