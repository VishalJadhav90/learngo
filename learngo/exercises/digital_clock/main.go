package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	saperator := [...]string{
		"░░░",
		"░█░",
		"░░░",
		"░█░",
		"░░░",
	}

	space := [...]string{
		"░░░",
		"░░░",
		"░░░",
		"░░░",
		"░░░",
	}

	digitsPattern := [...][5]string{
		{
			"███",
			"█░█",
			"█░█",
			"█░█",
			"███",
		},
		{
			"██░",
			"░█░",
			"░█░",
			"░█░",
			"███",
		},
		{
			"███",
			"░░█",
			"███",
			"█░░",
			"███",
		},
		{
			"███",
			"░░█",
			"███",
			"░░█",
			"███",
		},
		{
			"█░█",
			"█░█",
			"███",
			"░░█",
			"░░█",
		},
		{
			"███",
			"█░░",
			"███",
			"░░█",
			"███",
		},
		{
			"███",
			"█░░",
			"███",
			"█░█",
			"███",
		},
		{
			"███",
			"░░█",
			"░░█",
			"░░█",
			"░░█",
		},
		{
			"███",
			"█░█",
			"███",
			"█░█",
			"███",
		},
		{
			"███",
			"█░█",
			"███",
			"░░█",
			"███",
		},
	}

	// for i := 0; i < 5; i++ {
	// 	for j := range digitsPattern {
	// 		fmt.Printf("%s", digitsPattern[j][i])
	// 		fmt.Printf("%s", saperator[i])
	// 	}
	// 	fmt.Println()
	// }

	for {
		time.Sleep(time.Second)
		screen.Clear()
		screen.MoveTopLeft()
		now := time.Now()
		timeArray := [...]int{now.Hour(), now.Minute(), now.Second()}
		var digitsNums [6]int

		var counter int = 0
		for i := range timeArray {
			digitsNums[counter] = timeArray[i] / 10
			counter++
			digitsNums[counter] = timeArray[i] % 10
			counter++
		}
		//fmt.Printf("%v\n", digitsNums)
		for i := 0; i < 5; i++ {
			for j := range digitsNums {
				fmt.Printf("%s", digitsPattern[digitsNums[j]][i])
				if j != 5 {
					fmt.Printf("%s", space[i])
				}
				if j%2 == 1 && j != 5 {
					fmt.Printf("%s%s", saperator[i], space[i])
				}
			}
			fmt.Println()
		}
	}

}
