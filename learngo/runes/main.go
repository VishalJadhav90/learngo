package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "☁☂☔☠☣"
	bites := []byte(str)
	for _, b := range bites {
		fmt.Printf("%d\n", b)
	}
	fmt.Println("len of bites..", len(bites))
	fmt.Println("rune count in bites array..", utf8.RuneCount(bites))
	r := '☣'
	fmt.Printf("type of r is %T\n", r)
	newStr := string(bites)
	fmt.Printf("new string is :%s\n", newStr)
	fmt.Println("len of newString..", len(newStr))
	fmt.Println("rune count in newStr..", utf8.RuneCountInString(newStr))

	for i, run := range newStr {
		fmt.Printf("i: %d %x rune: %c\n", i, run, run)
	}
}
