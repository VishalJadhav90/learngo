package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

type collection []string

func main() {
	s.PrintElementAddr = true
	data := collection{"slices", "are", "awesome", "period"}
	change(data)
	s.Show("main data", data)
	fmt.Printf("main data slice header address: %p\n", &data)
}

func change(data collection) {
	data[2] = "brilliant"
	s.Show("change data", data)
	fmt.Printf("change data slice header address: %p\n", &data)
}
