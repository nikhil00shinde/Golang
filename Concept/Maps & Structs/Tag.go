package main

import (
	"fmt"
	"reflect" //to get tag of field 
)

type Animal struct{
	Name string `required max:"100"`
	Origin string
}

func main()  {

	t := reflect.TypeOf(Animal{})

	field,_ := t.FieldByName("Name")

	fmt.Println(field.Tag)
}