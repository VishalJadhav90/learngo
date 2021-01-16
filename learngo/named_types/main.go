package main

import (
	"fmt"
	"time"
)

type gram float32
type ounce float32

func main() {
	h, _ := time.ParseDuration("4h30m")
	fmt.Printf("type of h: %T", h)
	fmt.Println("Hours are..", h.Hours())
	addH := 4
	h = h * time.Duration(addH)
	fmt.Println("Hours are..", h.Hours())
	fmt.Println("nanosecs are ...", int64(h))

	var g gram = 1000
	var o ounce

	o = ounce(g) * 0.035274
	fmt.Printf("%g grams is %.2f ounces", g, o)
}
