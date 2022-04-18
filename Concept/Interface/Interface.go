package main

import (
	"fmt"
)
// any type that have method associated with it can implement a interface
// And any type can have method associated with it

func main(){
    var w Writer = ConsoleWriter{}
		w.Write([]byte("Hello World"))
}

//inside interface -> describing behaviour
type Writer interface {
    Write([]byte) (int,error) 
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n,err
}

// Name the interface with the method name + er