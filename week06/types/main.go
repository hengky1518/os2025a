package main

import (
	"fmt"
	"reflect"
)

func main() {
	//name:=2.71
	var name float64
	name = 2.71

	fmt.Println(name, reflect.TypeOf(name))

}
