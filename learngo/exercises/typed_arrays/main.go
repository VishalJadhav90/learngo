package main

import "fmt"

func main() {
	var bookcase [3]int{3,6,9}

	type cabinet [3]int
	cbnt := cabinet{3, 6, 9}

	fmt.Printf("Are they equal : %t", bookcase == cbnt)
}
