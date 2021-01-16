package main

import (
	"fmt"
)

type Duration int64

func main() {
	var d Duration = 100
	var i int64 = 200

	fmt.Println("vishal.... ", d, i)
	d, i = Duration(i), int64(d)
	fmt.Println("vishal.... ", d, i)

}
