package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println("Hello World")
	// var (
	// 	sum int
	// 	i   int
	// )
	// for {
	// 	if i > 5 {
	// 		break
	// 	}

	// 	if i%2 != 0 {
	// 		i++
	// 		continue
	// 	}

	// 	sum += i
	// 	fmt.Printf("sum: %02d i: %d\n", sum, i)
	// 	i++
	// }
	// fmt.Printf("final sum: %02d\n", sum)

	// fmt.Printf("%5s", "X")
	// for i := 1; i <= 5; i++ {
	// 	fmt.Printf("%5d", i)
	// }
	// fmt.Println()
	// for i := 1; i <= 5; i++ {
	// 	fmt.Printf("%5d", i)
	// 	for j := 1; j <= 5; j++ {
	// 		fmt.Printf("%5d", j*i)
	// 	}
	// 	fmt.Println()
	// }

	words := strings.Fields("lazy cat jumps again and again and again")

	for i := 0; i < len(words); i++ {
		fmt.Printf("#%-2d %q\n", i+1, words[i])
	}

	fmt.Println("............................")
	for i := range os.Args[1:] {
		fmt.Println(i)
	}

	var sum = 0
	var str = ""
	b, _ := strconv.ParseInt(os.Args[1], 10, 64)
	e, _ := strconv.ParseInt(os.Args[2], 10, 64)
	for ; b <= e; b++ {
		if b%2 == 1 {
			continue
		}
		sum = sum + int(b)
		str = str + " + " + strconv.FormatInt(int64(b), 10)
	}
	fmt.Printf("%s = %d\n", str[3:], sum)
}
