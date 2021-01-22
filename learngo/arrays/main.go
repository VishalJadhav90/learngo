package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	var abc [3]string
	fmt.Println("abc", abc)

	books := [...]string{
		"The Wolf Hall",
		"The Body Mass",
		"Gone with the wind",
		"Gentleman in Moscow",
	}

	fmt.Printf("books: %T\n", books)
	fmt.Printf("books: %q\n", books)
	fmt.Printf("books: %#v\n", books)

	var (
		wBooks [winter]string
		sBooks [summer]string
	)

	wBooks[0] = books[0]
	wBooks[0] = "Circee"

	for i := range sBooks {
		sBooks[i] = books[i+1]
	}

	fmt.Printf("wBooks: %#v\n", wBooks)
	fmt.Printf("sBooks: %#v\n", sBooks)
	fmt.Printf("books: %#v\n", books)

	var published [len(books)]bool
	published[0] = true
	published[len(books)-1] = true

	fmt.Printf("published: %#v\n", published)
}
