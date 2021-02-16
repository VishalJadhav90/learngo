package main

import "fmt"

func main() {
	var friends = []string{"amruta", "prasad", "pritam"}
	var distances = []int64{200000, 150000, 900000}
	var data = []uint8{'a', 'm', 'r', 'u', 't', 'a'}
	var ratios = []float64{2.3, 45.2, 3.9}
	var alives = []bool{true, false}
	fmt.Printf("%#v Is Nil: %t\n", friends, friends == nil)
	fmt.Printf("%#v Is Nil: %t\n", distances, distances == nil)
	fmt.Printf("%#v Is Nil: %t\n", data, data == nil)
	fmt.Printf("%#v Is Nil: %t\n", ratios, ratios == nil)
	fmt.Printf("%#v Is Nil: %t\n", alives, alives == nil)
}
