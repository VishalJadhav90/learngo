package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		num, err := strconv.ParseInt(arg, 10, 64)
		//fmt.Printf("#%d arg: %d\n", i, num)
		var flag int = 0
		if err == nil {
			for n := 2; int64(n) < num; n++ {
				if num%int64(n) == 0 {
					fmt.Println(num, "is not a prime")
					flag = 1
					break
				}
			}
			if flag == 0 {
				fmt.Println(num, "is a prime")
			}
		}

	}
}
