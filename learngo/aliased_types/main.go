package main

import (
	"fmt"
)

func main() {
	var (
		byteVal  byte
		uint8Val uint8
		intVal   int
	)

	type myType = uint8

	var v myType

	fmt.Printf("variable v: %d\n", v)

	uint8Val = byteVal

	var (
		runeVal  rune
		int32Val int32
	)

	runeVal = int32Val

	fmt.Printf("byteVal: %d uint8Val: %d intVal: %d runeVal: %d int32Val: %d", byteVal, uint8Val, intVal, runeVal, int32Val)
}
