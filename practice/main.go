package main

import (
	"fmt"
	"strconv"
)

func hammingWeight(num uint32) int {
	strRepr := strconv.FormatInt(int64(num), 2)
	fmt.Println(strRepr)
	count := 0
	for i := range strRepr {
		if strRepr[i] == '1' {
			count = count + 1
		}
	}
	return count
}

func main() {
	ones := hammingWeight(36)
	fmt.Println(ones)
}
