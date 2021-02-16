package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("")
	var inputs []int
	if len(os.Args) < 2 {
		fmt.Printf("Please input numbers..\n")
		return
	}
	for i := range os.Args {
		if i == 0 {
			continue
		}
		num, err := strconv.Atoi(os.Args[i])
		//fmt.Printf("num...%d\n", num)
		if err == nil {
			inputs = append(inputs, num)
		}
	}
	fmt.Printf("Input: %#v\n", inputs)
	sort.Ints(inputs)
	fmt.Printf("Sorted Input: %#v\n", inputs)
}
