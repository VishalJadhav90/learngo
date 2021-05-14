package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = lines + 1
	}
	fmt.Printf("lines are %d\n", lines)

	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR: ", err)
	}
}
