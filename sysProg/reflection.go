package main

import (
	"fmt"
	"log"
	"reflect"
)

type aStructure struct {
	A1 int
	A2 float64
	A3 string
}


func main() {

	type t1 int
	type t2 int

	x1 := t1(1)
	x2 := t2(2)
	x3 := 2

	x4 := aStructure{3,4,"A Structure"}

	st1 := reflect.ValueOf(&x1).Elem()
	st2 := reflect.ValueOf(&x2).Elem()
	st3 := reflect.ValueOf(&x3).Elem()
	st4 := reflect.ValueOf(&x4).Elem()

	fmt.Println("X1 Type: ",st1.Type())
	fmt.Println("X2 Type: ",st2.Type())
	fmt.Println("X3 Type: ",st3.Type())
	fmt.Println("X4 Type: ",st4.Type())

	typeOfX4 := st4.Type()

	log.Printf("X4 Type: %s\n",typeOfX4)
	log.Printf("The fields of %s are: \n",typeOfX4)

	for i := 0; i < st4.NumField(); i++ {
		fmt.Printf("%d: Field Name: %s ", i, typeOfX4.Field(i).Name)
		fmt.Printf("Type: %s ", st4.Field(i).Type())
		fmt.Printf("Value: %v\n",st4.Field(i).Interface())
	}

}
