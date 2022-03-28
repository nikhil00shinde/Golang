package main

import "fmt"

func main(){
	a := 8 //0001000 // 2^3

	fmt.Println(a << 3)//1000000 left bit shifting  // 2^3 * 2^3 = 2^6
	fmt.Println(a >> 3)//0000001 rigth bit shifting // 2^3 / 2^3 = 2^0

}