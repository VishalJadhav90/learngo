package main

import (
	"fmt"

	weights "github.com/VishalJadhav90/learngo/named_type_hierarchy/weigths"
)

func main() {

	type (
		Gram     int
		Kilogram Gram
		Ton      Kilogram
	)

	var (
		salt   Gram     = 100
		apples Kilogram = 5
		truck  Ton      = 10
	)
	fmt.Printf("salt: %d apples: %d truck: %d\n", salt, apples, truck)

	salt = Gram(apples)
	apples = Kilogram(truck)
	truck = Ton(apples)
	fmt.Printf("salt: %d apples: %d truck: %d\n", salt, apples, truck)

	salt = Gram(weights.Gram(100))

	fmt.Printf("Type of weights.Gram: %T\n", weights.Gram(1))
	fmt.Printf("Type of main.Gram: %T\n", Gram(1))

	type myGram weights.Gram
	fmt.Printf("Type of myGram: %T\n", myGram(1))
}
