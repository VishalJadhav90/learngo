package main

import (
	"fmt"
	"os"
)

func main() {

	switch city := os.Args[1]; city {
	case "Paris", "Lyon":
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	default:
		fmt.Println("Where?")
	}

}
