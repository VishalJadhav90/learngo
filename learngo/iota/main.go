package main

import "fmt"

func main() {
	const (
		a = iota
		_
		b
	)
	fmt.Println(a, b)
	const (
		sunday = iota
		monday
		tuesday
		wednsday
		thursday
		friday
		saturday
	)

	fmt.Println(sunday, monday, tuesday, wednsday, thursday, friday, saturday)

}
