package main

import "fmt"

func main() {
	subjects := []string{"GO", "Javascript", "Python", "Linux"}

	subjectsSlice := subjects[1:3]
	//subjects[0] = "Java"
	subjectsSlice[0] = "Java"
	subjectsSlice = append(subjectsSlice, "Go", "Kotlin", "Database")
	for _, subjects := range subjects {
		fmt.Println(subjects)
	}
	fmt.Println("==========")
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}
