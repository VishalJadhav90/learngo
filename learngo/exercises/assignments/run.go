package main

import "fmt"

func main() {
	color := "red"
	fmt.Println("go went gone", color)
	color = "blue"
	fmt.Println("go went gone", color)
	newcolor := color
	fmt.Println("new color is", newcolor)
	expColor := "reddish" + "green"
	fmt.Println("expColor is", expColor)
	var (
		width  = 100
		bredth = 50
	)
	area := width * bredth
	fmt.Println(area)

	lang, version := "go", 2
	fmt.Println(lang, version)

	first, second := multi()
	fmt.Println(first, second)
}

func multi() (int, int) {
	return 5, 4
}
