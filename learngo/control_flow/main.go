package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello world")

	n := 3
	if n > 5 {
		fmt.Printf("%d is less than %d", n, 5)
	} else {
		fmt.Printf("%d is greater than or equal to %d\n", n, 5)
	}

	fmt.Println(strings.Repeat("-", 25))

	speedB := 150.5
	faster := speedB > 100
	fmt.Println("is the other car going faster..", faster)

	fmt.Println("value is ....", !false)

	// if n {
	// 	fmt.Println("tis words")
	// } // expression in if condition has to be Boolean
}
