package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please give me a word to search\n")
		return
	}

	words := strings.Fields(corpus)

	for i := range words {
		if strings.Contains(strings.ToLower(words[i]), strings.ToLower(os.Args[1])) {
			fmt.Printf("#%-3d %s\n", i, words[i])
		}
	}
}
