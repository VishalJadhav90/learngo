package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	if len(os.Args) < 2 {
		fmt.Println("Tell me a book title")
	} else {
		found := false
		searchBook := strings.ToLower(os.Args[1])
		for i := range books {
			if strings.Contains(strings.ToLower(books[i]), searchBook) {
				fmt.Printf("+ %s\n", books[i])
				found = true
			}
		}
		if !found {
			fmt.Printf("We dont have book: %s\n", searchBook)
		}
	}
}
