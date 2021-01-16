package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	var s string
	s = "how are you?"
	s = `how are you?`
	fmt.Println(s)

	s = `
<html>
	<body>"Hello"</body>
</html>
	`
	fmt.Println(s)

	eq := "1 + 2 = "
	sum := 1 + 2
	fmt.Println(eq + strconv.Itoa(sum))

	fmt.Println(strconv.FormatBool(true))

	fmt.Println("size of a string...", len(s))

	c := "çok güzel"
	fmt.Println("special string len is...", len(c))

	fmt.Println("characters in rune is...", utf8.RuneCountInString(c))

}
