package main

import (
	"fmt"
)

func main() {
	speed := 50
	force := 2.5
	speed = int(float64(speed) * force)
	fmt.Println(speed)

	var apples int
	var oranges int32

	apples = int(oranges)
	fmt.Println(apples)

	char := string(65)
	fmt.Println(char)

	char = string([]byte{104, 105})
	fmt.Println(char)

	area := 10.5
	div := 2
	fmt.Println(area / float64(div))
}
