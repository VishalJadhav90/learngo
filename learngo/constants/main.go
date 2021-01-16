package main

import "fmt"

func main() {
	const meters int = 100
	const factor int = 0

	cm := 100
	m := cm / meters
	fmt.Printf("%dcm is %dm\n", cm, m)

	cm = 200
	m = cm / meters
	fmt.Printf("%dcm is %dm\n", cm, m)

	//fmt.Println("CompileTimeError", meters/factor)

	// typeless constants
	// const min = 1 + 1
	// const pi = 3.14 + min
	// const version = "2.0.1" + "-beta"
	// const debug = !false
	// fmt.Printf("%d, %g, %s, %t\n", min, pi, version, debug)

	// var a, b, c int = 10, 20, 30
	// a = b
	// b = c
	// c = a
	// fmt.Println(a, b, c)

	// const (
	// 	min int = 1
	// 	max
	// )

	// fmt.Println(min, max)

	const min = 42
	var f float64 = min
	var b byte = min
	var r rune = min
	fmt.Println("Variable F is", f, "B is", b, "R is", r)

}
