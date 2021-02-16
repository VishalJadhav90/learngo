package main

import (
	"sort"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	agesArray := [3]int{35, 15, 25}
	ages := agesArray[:]
	ages[0] = 100
	ages[1] = 50
	s.Show("agesArray", agesArray)
	s.Show("ages", ages)

	grades := [...]float64{40, 10, 20, 50, 60, 70}

	newGrades := append([]float64{}, grades[:]...)
	front := newGrades[:3]

	sort.Float64s(front)

	s.PrintBacking = true
	s.MaxPerLine = 7
	s.Show("grades", grades[:])
	s.Show("front", front)
	s.Show("end", newGrades[4:])
	s.Show("newGrades", newGrades)
}
