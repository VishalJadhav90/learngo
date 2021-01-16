package main

import (
	"fmt"
)

func main() {
	/*
		const (
			feetToMeters = 0.3048
			feetToYards  = 0.3333
		)

		feet, _ := strconv.ParseFloat(os.Args[1], 64)
		meters := math.Round(feet * feetToMeters)
		yards := math.Round(feet * feetToYards)
		fmt.Printf("%g feet is %g meters\n", feet, meters)
		fmt.Printf("%g feet is %g yards\n", feet, yards)
	*/

	const minsPerDay = 24 * 60
	const weekDays = 7
	fmt.Println("number of mins in two weeks", weekDays*minsPerDay*2)

	const hoursPerDay = 34
	const daysPerWeek = 7
	fmt.Printf("number of hours %d in %d weeks\n", 5*daysPerWeek*hoursPerDay, 5)
}
