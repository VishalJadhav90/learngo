package main

import (
	"fmt"
	"strconv"
	"strings"
)

type rentDetails struct {
	city  string
	size  int
	beds  int
	baths int
	price int
}

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var details []rentDetails
	var headers []string = strings.Split(header, separator)
	for _, det := range strings.Split(data, "\n") {
		//fmt.Printf("det...%#v\n", det)
		splitdet := strings.Split(det, separator)
		//fmt.Printf("splitdet...%#v\n", splitdet)
		city := splitdet[0]
		size, _ := strconv.Atoi(splitdet[1])
		beds, _ := strconv.Atoi(splitdet[2])
		baths, _ := strconv.Atoi(splitdet[3])
		price, _ := strconv.Atoi(splitdet[4])
		details = append(details, rentDetails{city: city, size: size, beds: beds, baths: baths, price: price})
	}

	for i := range headers {
		fmt.Printf("%-20s", headers[i])
	}
	fmt.Printf("\n===========================================================================================\n")

	for i := range details {
		fmt.Printf("%-20s%-20d%-20d%-20d%-20d\n", details[i].city, details[i].size,
			details[i].beds, details[i].baths, details[i].price)
	}

}
