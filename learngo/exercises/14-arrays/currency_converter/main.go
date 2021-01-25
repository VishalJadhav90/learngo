package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	const (
		EUR = iota
		GBP
		JPY
	)

	ratios := [...]float64{
		EUR: 0.88,
		GBP: 0.78,
		JPY: 113.02,
	}

	cur := [...]string{
		EUR: "EUR",
		GBP: "GBP",
		JPY: "JPY",
	}

	if len(os.Args) != 2 {
		fmt.Printf("Please provide the amount to be converted\n")
		return
	}

	dollars, err := strconv.ParseFloat(os.Args[1], 32)
	if err == nil {
		for i := range ratios {
			fmt.Printf("%.3g USD is %.3g %s\n", dollars, dollars*ratios[i], cur[i])
		}
	}

}
