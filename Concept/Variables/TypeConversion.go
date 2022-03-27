package main

import (
	"fmt"
	"strconv"
)

func main(){

	//integer to float
	// var i int = 42
	// fmt.Printf("%v, %T\n",i,i)

	// var j float32
	// j = float32(i)

	// fmt.Printf("%v, %T\n",j,j)

//----------------------------------------------

	// float to integer
	// var i float32 = 42
	// fmt.Printf("%v, %T\n",i,i)

	// var j int
	// j = int(i)

	// fmt.Printf("%v, %T\n",j,j)

//-----------------------------------------------

  // integer to string
	var i int = 42
	fmt.Printf("%v, %T\n",i,i)

	var j string
	// j = string(i)  //42 unicode of * (ascii)
	j = strconv.Itoa(i)

	fmt.Printf("%v, %T\n",j,j)




}