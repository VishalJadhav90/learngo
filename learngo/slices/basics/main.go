package main

import "fmt"

func main() {
	fmt.Println()
	var slc = []int{}
	fmt.Printf("%#v\n", slc)
	fmt.Printf("Empty? %t\n", slc == nil)
	fmt.Printf("Empty? %t\n", len(slc) == 0)

	var books [5]string
	books[0] = "dracula"
	books[1] = "island"
	books[2] = "sims"

	newBooks := books
	books[3] = "blahblah"

	fmt.Printf("books: %#v\n", books)
	fmt.Printf("newBooks: %#v\n", newBooks)

	games := []string{"kokemon", "sims", "doremon", "pikachu"}
	fmt.Printf("games: %#v\n", games)
	newGames := []string{"pacman", "sone"}

	newGames = games
	games = nil
	fmt.Printf("games: %#v\n", games)
	fmt.Printf("newGames: %#v\n", newGames)

}
