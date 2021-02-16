package main

import "fmt"

func main() {
	toppings := []string{"black olives", "green pappers"}
	var pizza []string
	toppings = append(toppings, "onions")
	toppings = append(toppings, "extra cheese")
	pizza = append(pizza, toppings...)
	fmt.Printf("pizza	:%s\n", pizza)
}
