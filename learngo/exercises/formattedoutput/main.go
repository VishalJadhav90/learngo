package main

import (
	"fmt"
	"os"
)

func main() {
	name, lastname, age := os.Args[1], os.Args[2], os.Args[3]
	fmt.Printf("my name is %s surname is %q age is %v\n", name, lastname, age)

	claim := false
	fmt.Printf("These are %t claims\n", claim)

	temp := 29.5
	fmt.Printf("Temparature today is %.2f degrees\n", temp)

	fmt.Printf("Type of temp is %T\n", temp)
}
