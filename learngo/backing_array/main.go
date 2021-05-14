package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true
	nums := []int{}
	fmt.Printf("%#v\n", nums)
	var nums1 []int
	fmt.Printf("%#v\n", nums1)
	s.Show("no backing array", nums)
	nums = append(nums, 1, 3)
	s.Show("allocates", nums)
	nums = append(nums, 2)
	s.Show("allocates", nums)
	nums = append(nums, 4)
	s.Show("allocates", nums)

	nums = append(nums, nums[2:]...)
	s.Show("allocates", nums)

	nums = append(nums[:2], 7, 9)
	s.Show("allocates", nums)

	nums = nums[:6]
	s.Show("allocates", nums)
}
