package main

import "fmt"

func main() {
	type EngChar byte
	var c EngChar = 'a'
	fmt.Printf("Named data types: %c\n", c)

	type SpChar rune
	var sc SpChar = 'ä'
	fmt.Printf("Special Char: %c\n", sc)

	type Year uint16
	var bYear Year = 1990
	fmt.Printf("Birth Year: %d\n", bYear)

	type Month uint8
	var bMonth Month = 7
	fmt.Printf("Birth Month: %d\n", bMonth)

	type Speed uint64
	var lSpeed Speed = 3409345
	fmt.Printf("Speed of light: %d Kms\n", lSpeed)

	type Angle float32
	var cAngle Angle = 12.34
	fmt.Printf("Angle of Arc: %g\n", cAngle)

	var (
		width  uint8
		height uint16
	)

	width, height = 255, 265
	width += 10
	fmt.Printf("Width: %d Height: %d\n", width, height)
	fmt.Printf("are they equal: %t\n", uint16(width) == height)
}
