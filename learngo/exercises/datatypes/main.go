package main

import "fmt"

var glob int

func main() {
	fmt.Println(-100, 100, .10, 5.0, true, false, "", "vishal", 0x12)
	var (
		myInt    int
		myBool   bool
		myFloat  float64
		myString string
	)
	_ = myInt
	_ = myBool
	_ = myFloat
	_ = myString

	var (
		myInt8  int8
		myInt16 int16
	)

	var v rune
	_ = v
	var c byte
	c = 'a'
	_ = c
	const r = `⌘`
	_ = myInt8
	_ = myInt16
	_ = r

	safe, notsafe := true, false
	_, _ = safe, notsafe
}
