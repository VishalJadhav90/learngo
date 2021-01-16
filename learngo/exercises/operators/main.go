package main

import "fmt"

func main() {
	fmt.Printf("%d %g %g %d %d\n", 50+25, 50-2.5, 50*0.5, 25/3, -(5 + 2))
	fmt.Printf("%g\n", 5.0/2)

	counter, factor := 45, 0.5
	counter *= 5
	factor /= 2
	fmt.Printf("counter: %d factor: %g\n", counter, factor)

	cmt := 1
	cmt++
	cmt--
	cmt += 5
	cmt *= 10
	cmt /= 5
	fmt.Printf("cmt %d\n", cmt)

	var (
		radius = 10.0
		area   float64
	)
	area = 22.0 / 7 * float64(radius) * float64(radius)
	fmt.Printf("radius: %g -> area: %g\n", radius, area)

	surfaceArea := 4 * 22.0 / 7.0 * radius * radius
	fmt.Printf("radius: %g -> surface_area: %.2f\n", radius, surfaceArea)
}
