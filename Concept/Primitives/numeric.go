package main

import "fmt"

func main(){
	// var n int = 42  //int8(-128 to +128) int16 int32 int64

	var n uint16 = 42  //uint8 uint16 uint32
	fmt.Printf("%v %T\n",n,n)
}