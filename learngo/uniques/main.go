package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	max, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		return
	}
	var uniques []int
loop:
	for len(uniques) < int(max) {
		num := rand.Intn(int(max)) + 1
		for _, u := range uniques {
			if u == num {
				continue loop
			}
		}
		//uniques[i] = num
		uniques = append(uniques, num)
	}

	fmt.Printf("random uniques %#v\n", uniques)

	sort.Ints(uniques)
	fmt.Printf("sorted uniques %#v\n", uniques)

	var num = [5]int{4, 5, 2, 8, 0}
	sort.Ints(num[:])
	fmt.Printf("sorted num %#v\n", num)

}
