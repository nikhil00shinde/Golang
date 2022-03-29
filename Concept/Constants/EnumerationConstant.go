package main

import "fmt"

const (
	a = iota //iota -> it's scope is within this constant block
	b = iota
	c
)

const (
	a2 = iota
)

const (
	_ = iota + 5
	catSpecialist
	dogSpecialist
)

const (
	_ = iota
	KB = 1 << (10 * iota) //for next expression will get evaluated
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main()  {
	// fmt.Printf("%v\n",a)
	// fmt.Printf("%v\n",b)
	// fmt.Printf("%v\n",c)
	// fmt.Printf("%v\n\n",a2)

	// fmt.Printf("%v\n",catSpecialist)
	// fmt.Printf("%v\n",dogSpecialist)
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n",fileSize/GB)
	fmt.Printf("%v\n",GB)


}

// Example
// package main

// import "fmt"

// const (
// 	errorSpecialist = iota
// 	catSpecialist
// 	dogSpecialist
// 	snakeSpecialist
// )


// func main(){
//   var specialistType int
// 	fmt.Printf("%v\n",specialistType == errorSpecialist)
// }













