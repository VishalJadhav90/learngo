package main

import (
	"fmt"
	"path"
	"time"
)

func main() {
	name, age := "vishal", 31
	sum := 25 + 8
	fmt.Println(name, age, sum)
	_ = sum
	salary := 25000
	salary, experience := 60000, 5.5
	fmt.Println(salary, experience)

	var (
		speed int
		now   time.Time
	)

	speed, now = 100, time.Now()
	fmt.Println(speed, now)

	var (
		curRate  = 10.5
		prevRate = 12.5
	)
	curRate, prevRate = prevRate, curRate
	fmt.Println(curRate, prevRate)

	dir, file := path.Split("google.com/songs")
	fmt.Println(dir, file)
}
