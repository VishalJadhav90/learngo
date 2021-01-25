package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 6 {
		fmt.Println("Program supports only 5 arguments")
		return
	}
	var inputArray [5]int64
	for i := range os.Args {
		if i == 0 {
			continue
		}
		num, err := strconv.ParseInt(os.Args[i], 10, 64)
		if err == nil {
			inputArray[i-1] = num
		}
	}
	fmt.Printf("Input array is %v\n", inputArray)
	for i := 1; i < len(inputArray); i++ {
		for j := 0; j < len(inputArray)-i; j++ {
			if inputArray[j] > inputArray[j+1] {
				temp := inputArray[j]
				inputArray[j] = inputArray[j+1]
				inputArray[j+1] = temp
			}
		}
	}
	fmt.Printf("Sorted array is %v\n", inputArray)
}
