package main

import (
	"fmt"
	"runtime"
)

/*
function_main will be an entry point for the execution of the program
*/
func main() {
	fmt.Println("hello there...")
	fmt.Println("I am here")
	fmt.Println(runtime.Version())
}
