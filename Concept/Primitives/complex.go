package main

import "fmt"

func main(){
	// float32 float64
//  var n complex64 = 1 + 2i  //complex64 complex128 (float64/32 + float64/32*i)
  var n complex128 = complex(5,12)
	fmt.Printf("%v %T\n\n",n,n)
  
	// Get real & imaginery part
	fmt.Printf("%v %T\n",real(n),real(n))
	fmt.Printf("%v %T\n",imag(n),imag(n))


	// a := 1 + 2i
	// b := 2 + 5.2i
  // fmt.Println(a + b)
  // fmt.Println(a - b)
  // fmt.Println(a * b)
  // fmt.Println(a / b)

}