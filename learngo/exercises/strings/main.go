package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	path := `c:\program files\duper super\fun.txt
c:\program files\really\funny.png `
	fmt.Println(path)
	name := "İNANÇ"
	fmt.Println("char count is..", utf8.RuneCountInString(name))
	fmt.Println(name + strings.Repeat("!", utf8.RuneCountInString(name)))
	uppercase := "VISHAL"
	fmt.Println("lowercase is ", strings.ToLower(uppercase))
	msg := "inanç           "
	fmt.Println("msg was", msg, "trimmed msg\n"+strings.TrimRight(msg, " ")+"end")
}
