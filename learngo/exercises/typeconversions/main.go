package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a, b := 10, 5.5
	fmt.Println(float64(a) + b)

	celcius, _ := strconv.ParseFloat(os.Args[1], 64)
	far := celcius*1.8 + float64(32)
	fmt.Printf("celcius: %g = far: %g\n", celcius, far)
}
