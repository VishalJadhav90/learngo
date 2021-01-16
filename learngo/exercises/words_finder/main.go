package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words := os.Args[1:]
	stmt := "cat jumps again and again and again"

queries:
	for _, word := range words {
		word = strings.ToLower(word)
	search:
		for j, field := range strings.Fields(stmt) {
			switch word {
			case "and", "or":
				break search
			}
			if word == field {
				fmt.Printf("found %q at index:%d\n", word, j+1)
				continue queries
			}
		}
	}
}
