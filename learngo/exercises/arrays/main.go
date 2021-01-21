package main

import (
	"fmt"
	"strings"
)

func main() {
	var names = [...]string{"srp", "prasad", "piyush"}
	var distance = [...]int{100, 200, 300}
	var buff = [5]byte{1, 2, 4}
	var exchratio = [...]float32{69.0, 71.3, 69.7}
	var status = [4]bool{true, false, true, true}
	var zero = []byte{}

	// fmt.Printf("names type: %T\n", names)
	// fmt.Printf("distance type: %T\n", distance)
	// fmt.Printf("buff type: %T\n", buff)
	// fmt.Printf("exchratio type: %T\n", exchratio)
	// fmt.Printf("status type: %T\n", status)
	// fmt.Printf("zero type: %T\n", zero)

	// fmt.Println()
	// fmt.Printf("names type&value: %#v\n", names)
	// fmt.Printf("distance type&value: %#v\n", distance)
	// fmt.Printf("buff type&value: %#v\n", buff)
	// fmt.Printf("exchratio type&value: %#v\n", exchratio)
	// fmt.Printf("status type&value: %#v\n", status)
	// fmt.Printf("zero type&value: %#v\n", zero)

	// fmt.Println()
	// fmt.Printf("names value: %v\n", names)
	// fmt.Printf("distance value: %v\n", distance)
	// fmt.Printf("buff value: %v\n", buff)
	// fmt.Printf("exchratio value: %v\n", exchratio)
	// fmt.Printf("status value: %v\n", status)
	// fmt.Printf("zero value: %v\n", zero)

	fmt.Printf("\nnames\n")
	fmt.Printf("==============================================\n")
	for i := range names {
		fmt.Printf("names[%d]: %v\n", i, names[i])
	}

	fmt.Printf("\ndistances\n")
	fmt.Printf("==============================================\n")
	for i := range distance {
		fmt.Printf("distances[%d]: %v\n", i, distance[i])
	}

	fmt.Printf("\ndata\n")
	fmt.Printf("==============================================\n")
	for i := range buff {
		fmt.Printf("data[%d]: %v\n", i, buff[i])
	}

	fmt.Printf("\nratios\n")
	fmt.Printf("==============================================\n")
	for i := range exchratio {
		fmt.Printf("ratio[%d]: %v\n", i, exchratio[i])
	}

	fmt.Printf("\nalives\n")
	fmt.Printf("==============================================\n")
	for i := range status {
		fmt.Printf("alives[%d]: %v\n", i, status[i])
	}

	fmt.Printf("\nzero\n")
	fmt.Printf("==============================================\n")
	for i := range zero {
		fmt.Printf("zero[%d]: %v\n", i, zero[i])
	}

	// fix below code snippet - use array literals
	fmt.Printf("\n")
	var bnames = [3]string{
		"Einstein",
		"Shepard",
		"Tesla",
	}

	var books = [5]string{
		"Kafka's Revenge",
		"Stay Golden",
	}

	fmt.Printf("%q\n", bnames)
	fmt.Printf("%q\n", books)

	// fix below code snippet - compare arrays
	week := [4]string{"Monday", "Tuesday"}
	wend := [4]string{"Saturday", "Sunday"}
	fmt.Println("is week not equal to wend...", week != wend)

	evens := [...]int{2, 4, 6, 8, 10}
	evens2 := [...]int{2, 4, 6, 8, 10}

	fmt.Println("is evens equal to evens2...", evens == evens2)

	// Use     : uint8 for one of the arrays instead of byte here.
	// Remember: Aliased types are the same types.
	image := [5]byte{'h', 'i'}
	var data [5]byte

	fmt.Println(data == image)

	fmt.Println("\n-------create a new array by copying existing array--------")
	var mybooks = [...]string{
		"kafka revenge",
		"stay golden",
		"everythingship",
	}

	upper := mybooks
	lower := mybooks

	for i := range upper {
		upper[i] = strings.ToUpper(upper[i])
	}

	for i := range lower {
		lower[i] = strings.ToLower(lower[i])
	}

	fmt.Println()
	fmt.Printf("upper: %#v\nlower: %#v\n", upper, lower)
}
