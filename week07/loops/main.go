package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var day int = now.Day()
	fmt.Println(day)
	// broken := "G# r#cks!"
	// replacer := strings.NewReplacer("#", "o")
	// fixed := replacer.Replace(broken)
	// fmt.Println(fixed)
}
