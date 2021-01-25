package main

import "fmt"

func main() {
	type (
		integer int

		bookcase [5]int
		cabinet  [5]int
	)
	blue := bookcase{6, 9, 3, 2, 1}
	reed := cabinet{6, 9, 3, 2, 1}
	fmt.Printf("Are they equal?\n")
	if cabinet(blue) == reed {
		fmt.Printf("blue and reed are equal\n")
	} else {
		fmt.Printf("blue and reed are not equal\n")
	}

	fmt.Printf("blue: %#v\n", blue)
	fmt.Printf("reed: %#v\n", reed)

	type rainbow [5]int
	whit := rainbow{6, 9, 3, 2, 1}
	fmt.Printf("whit: %#v\n", whit)
}
