package main

import (
	"fmt"
	"math"
)

// interface is really a way looking at a set of related objects or types
// let us defined some kind that is similar between objects or similar between types, so that we can use this types one in the same with their relative behaviour

type shape interface{
	area() float64
}
// anything any type or struct that has this area method and returns float64 is of type shape

type shape2 interface{
	perim() float64
}

type circle struct{
	radius float64
}

type rect struct{
	width float64
	height float64
}

// We add this method which accept the pointer, not just instance / copy
func (r *rect)  area() float64{
	return r.width * r.height
}

func (c circle)  area() float64{
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64{
	return s.area()
}

func main(){
	c1 := circle{4.5}
	r1  := rect{5,7}
  
	// shapes := []shape{c1,r1}
	shapes := []shape{c1,&r1}
  //  it is a good practice to always pass the pointer when you are making kind slices stuff like, soo that you need access to the pointer, you have it(IT DOESNOT HURT TO PASS THE POINTER)
	// if a method has a pointer on it, then u need to pass the pointer otherwise it won't implement the interface correctly

	// as soon as , I use interface as a type all of the behaviour that, I know is true about circle and rectangle, I cannot access through this interface
	  
	for _,shape := range shapes{
		fmt.Println(getArea(shape))
	}
}


