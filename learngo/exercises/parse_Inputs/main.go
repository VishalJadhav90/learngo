package main

import "fmt"

func main() {
	// val, _ := strconv.ParseInt(os.Args[1], 10, 8)
	// fmt.Printf("Number is %d\n", int8(val))

	// var av int64
	// av, _ = strconv.ParseInt(os.Args[2], 10, 16)
	// fmt.Printf("another Number %d\n", av)

	// bv, _ := strconv.ParseInt(os.Args[3], 2, 8)
	// fmt.Printf("Number converted from Binary Sequence %d\n", bv)

	// dur, _ := time.ParseDuration("1h30m")
	// add, _ := strconv.ParseInt(os.Args[1], 10, 64)
	// fmt.Printf("Old hour: %g\n", dur.Hours())

	// dur = dur * time.Duration(add)
	// fmt.Printf("New hour: %g\n", dur.Hours())

	// type Feet float64
	// type Meters float64

	// var (
	// 	feet   Feet
	// 	meters Meters
	// )

	// input, _ := strconv.ParseFloat(os.Args[1], 64)
	// feet = Feet(input)
	// meters = Meters(feet) * Meters(0.3048)
	// fmt.Printf("Feet: %g Meters: %g\n", feet, meters)

	type (
		Temperature float64
		Celcius     Temperature
		Fahrenheit  Temperature
	)

	var (
		celcius       Celcius     = 15.5
		fahr          Fahrenheit  = 59.9
		celciusDegree Temperature = 10.5
		fahrDegree    Temperature = 10.5
		factor                    = 2.
	)

	celcius *= Celcius(celciusDegree) * Celcius(factor)
	fahr *= Fahrenheit(fahrDegree) * Fahrenheit(factor)

	fmt.Println(celcius, fahr)

}
