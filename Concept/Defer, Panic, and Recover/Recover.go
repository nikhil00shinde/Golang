package main

import (
	"fmt"
	"log"
)

func main(){
	//defer statement -> also take function call
	fmt.Println("start")
	// defer fmt.Println("this was deferred")
  
	//defer statement -> also take function call
	defer func(){
		// anonumous function 
		// recover func -> return nil if application is panicking
		// if it isn't nil it actually return error then actually it causing the application to panic
		if err := recover(); err != nil{
			log.Println("Error:",err)
		}
	}()
	panic("something bad happened")
	fmt.Println("end")
}