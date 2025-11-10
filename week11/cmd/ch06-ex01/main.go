package main

import "fmt"

func main() {
	subjects := []string{"GO", "Javascript", "Python", "Linux"}

	subjectsSlice := subjects[1:3]

	for _, subjects := range subjects {
		fmt.Println(subjects)
	}

	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}
