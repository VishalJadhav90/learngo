package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	inputstr := os.Args[1]
	excl := strings.Repeat("!", len(inputstr))
	fmt.Println(strings.ToUpper(inputstr + excl))
	fmt.Println(utf8.RuneCountInString("hétérogénéité"))
}
