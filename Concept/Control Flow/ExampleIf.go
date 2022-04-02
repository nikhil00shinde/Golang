package main

import (
	"fmt"
	"math"
)

func main(){
  number := 50
	guess := 30

	if guess < 1 ||  guess > 100 {
		  fmt.Println("The guess must be greater than 1!")
	}else if guess > 100{
    fmt.Println("The guess must be less than 100")
	}else {
		  if guess >  number {
				fmt.Println("Too high")
			}
			if guess < number {
				fmt.Println("Too low")
			}
			if guess == number {
				fmt.Println("You got it!")
			}
			fmt.Println(number <= guess, number >= guess,number != guess)
	}

	fmt.Println(!returnTrue())

  myNum := 0.123

	// if myNum == math.Pow(math.Sqrt(myNum),2) {
	// 	fmt.Println("These are the same")
	// } else {
	// 	fmt.Println("These are 4 diffrent")
	// }

	if math.Abs(myNum / math.Pow(math.Sqrt(myNum),2) - 1) < 0.001{
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are 4 diffrent")
	}
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}


// concept -> SHORT CIRCUTING (if one part of OR(||) statement is true) then it will execute test, and will not evaluate other OR test