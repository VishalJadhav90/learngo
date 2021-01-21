package main

import "fmt"

func main() {
	var (
		blue = [3]int{6, 9, 3}
		red  = [3]int{6, 9, 3}
	)

	fmt.Println("Are they equal...", blue == red)

	var (
		arr1 = [...]int{1, 2, 3}    //size 3
		arr2 = [...]int{1, 2, 3, 4} //size 4
	)
	_, _ = arr1, arr2
	// fmt.Println("Are they equal...", arr1==arr2); // this comparison throws exception as type of array is not equal

}
