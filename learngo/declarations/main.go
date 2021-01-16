package main

import "fmt"

func main() {
	name, lastname := "vishal", "jadhav"
	_, _ = name, lastname
	fmt.Println(name, lastname)

	birth, death := 81323, 34343
	fmt.Println(birth, death)

	on, off := true, false
	fmt.Println(on, off)

	nfiles, valid, msg := 10, false, "hello"
	fmt.Println(nfiles, valid, msg)

	var safe bool
	safe, speed := false, 50
	fmt.Println(safe, speed)

	speed = 70

	// when you dont know initial value dont use short declaration
	age := 0 // this is not recommended
	_ = age

	// however if you know initial value then use short declaration
	myage := 31 // this is okay
	_ = myage
}
