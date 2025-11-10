package main

import "fmt"

func main() {
	subjects := []string{"GO", "", "Python"} //innitialized

	for _, subjects := range subjects {
		fmt.Println(subjects)
	}
}
