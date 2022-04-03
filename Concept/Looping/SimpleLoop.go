package main

import (
	"fmt"
)

func main(){
	//initializer condition increment/decrement  
	for i := 0; i < 5; i++ {
		//i is scoped to for loop
		fmt.Println(i)
	}

	// for i,j := 0,0; i < 5; i,j = i+1,j+2{
	// 	fmt.Println(i,j)
	// }


	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// 	if i%2 == 0 {
	// 		i /= 2
	// 	}else {
	// 		i = i*2 + 1
	// 	}
	// }

	// i := 0
	// for ;i < 5;i++{
	// 	fmt.Println(i)
	// }

	// i := 0
	// for ;i < 5;{
	// 	fmt.Println(i)
	// 	i++
	// }

	// i := 0
	// for i < 5{      //work like while loop
	// 	fmt.Println(i)
	// 	i++
	// }


	// i := 0  work like do-while loop
	// for {
	// 	fmt.Println(i)
	// 	i++
	// 	if i == 5{
	// 		break
	// 	}
	// }

	
	for i := 0;i < 10;i++{
		if(i%2==0){
			continue
		}
		fmt.Println(i)
		
	}
	
}