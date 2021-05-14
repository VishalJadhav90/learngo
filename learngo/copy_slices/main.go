package main

import (
	"fmt"
	"strconv"
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	data := []float64{10, 25, 30, 50}
	s.Show("Data...", data)

	newData := make([]float64, len(data))
	n := copy(newData, data)
	fmt.Println("copied elements are", n)
	s.Show("newData...", newData)

	//copy + make => append
	easyData := append([]float64{}, data...)
	s.Show("easyData...", easyData)

	spendings := fetch()

	s.Show("easyData...", spendings)

	total := 0
	for i, daily := range spendings {
		fmt.Printf("%d->%#v\n", i, daily)
		for _, exp := range daily {
			total += exp
		}
	}
	fmt.Printf("Total expenditure...%d\n", total)

	lyric := []string{"show", "me", "my", "silver", "lining"}
	part1 := lyric[1:3:5]
	s.Show("part1", part1)
	part := lyric[2:3:5]
	s.Show("part", part)
}

func fetch() [][]int {
	content := `200 100
25 10 45 60
5 15 35
95 10
50 25`
	spendings := make([][]int, 5, 5)
	for i, dailyStr := range strings.Split(content, "\n") {
		spendings[i] = make([]int, len(strings.Fields(dailyStr)))
		for j, spending := range strings.Split(dailyStr, " ") {
			spend, _ := strconv.Atoi(spending)
			spendings[i][j] = spend
		}
		fmt.Println("Calculated spendings for day: ", i)
	}
	return spendings
}
