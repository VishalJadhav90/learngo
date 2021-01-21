package main

import "fmt"

func main() {
	fmt.Println("")
	var (
		prev = [...]string{"The Wolf Hall", "Circee", "The Body Mass"}
	)

	books := prev

	for i := range prev {
		books[i] = prev[i] + " 2nd Ed."
	}

	fmt.Printf("prevBooks: %#v\n", prev)
	fmt.Printf("books: %#v\n", books)
}
