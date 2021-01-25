package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var numbers [5]int64

	var (
		sum   int64 = 0
		count int64 = 0
	)

	for i := range os.Args {
		if i == 0 {
			continue
		}

		num, err := strconv.ParseInt(os.Args[i], 10, 64)
		if err == nil {
			numbers[count] = num
			count++
			sum += num
		} else {
			numbers[count] = 0
			count++
		}
	}
	fmt.Printf("Your numbers: %v\n", numbers)
	fmt.Printf("Average: %g\n", float64(sum)/float64(count))

}
