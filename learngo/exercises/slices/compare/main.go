package main

import (
	"fmt"
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	namesA := "Da vinci,Woznik,Carmack"
	namesB := []string{"Woznik", "Da vinci", "Carmack"}
	names := strings.Split(namesA, ",")
	notEqual := false
	if len(names) != len(namesB) {
		notEqual = true
	} else {
		for i := range namesB {
			if names[i] != namesB[i] {
				notEqual = true
			}
		}
	}

	fmt.Printf("Are %#v & %#v EQUAL = %t\n", names, namesB, !notEqual)

	s.Show("names", names)

	var todo []string

	todo = append(todo, "sing")
	//todo = append(todo, 42)
	s.Show("todo", todo)

	todo = append(todo, "bath", "run")
	s.Show("todo", todo)

	tmrw := []string{"work", "play"}
	todo = append(todo, tmrw...)
	s.Show("todo", todo)
}
