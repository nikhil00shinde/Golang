// When I implementing interface. If I use a value type, the methods that implements the interface have to all have value reciever 
// If I'm implementing interface with a pointer then I just have to have methods their regardless of the reciever type
package main

import (
	"fmt"
)

func main(){
	 var wc WriterCloser = &myWriterCloser{}
	 fmt.Println(wc)
}

type WriterCloser interface{
   Writer
	 Closer	
}

type Writer interface{
	Write() (int,error)
}

type Closer interface{
	Close() error
}

type myWriterCloser struct{}

func (mwc myWriterCloser) Write() (int,error){
	return 0,nil
}

func (mwc myWriterCloser) Close() (error){
	return nil
}