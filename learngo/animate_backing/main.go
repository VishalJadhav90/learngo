package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true
	s.MaxPerLine = 30
	s.Width = 150

	// var nums []int
	// screen.Clear()
	// for cap(nums) <= 128 {
	// 	screen.MoveTopLeft()
	// 	s.Show("nums", nums)
	// 	nums = append(nums, rand.Intn(9)+1)
	// 	time.Sleep(time.Second / 4)
	// }

	words := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s.Show("words", words)
	words1 := append(words[:2], 22, 33)
	s.Show("words", words1)
	s.Show("words", words)
	newWords := words[2:5:9]
	s.Show("newWords", newWords)
}
