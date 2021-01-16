package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	const meterPerFeet = 0.3048
	feet, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid feet entered: ", os.Args[1])
	} else {
		fmt.Printf("%g feet equals to %g meters\n", feet, feet*meterPerFeet)
	}
}
