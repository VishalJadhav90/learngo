package main

import "fmt"

func main() {
	var bookcase = [3]int{3, 6, 9}

	type cabinet [3]int
	cbnt := cabinet{3, 6, 9}

	fmt.Printf("Are they equal : %t", bookcase == cbnt)

	type speeds [3]int
	carSpeed := speeds{55, 66, 78}
	_ = carSpeed
	type averages [3]int
	avgSpeed := averages{66, 76, 80}
	_ = avgSpeed

	//fmt.Printf("are the equal %t", carSpeed == avgSpeed) //this does not work as both are named types
}
