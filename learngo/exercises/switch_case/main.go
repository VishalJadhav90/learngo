package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	switch hour := time.Now().Hour(); true {
	case 0 < hour && hour < 12:
		fmt.Println("Good Morning")
	case 12 <= hour && hour < 16:
		fmt.Println("Good Evening")
	case 16 < hour && hour <= 24:
		fmt.Println("Good Night")
	}

	n := 8

	switch {
	case n > 5, n < 1:
		fmt.Println("n is big")
	case n == 8:
		fmt.Println("n is 8")
	}

	// if len(os.Args) != 2 {
	// 	fmt.Println("Give me the magnitude of the earthquake")
	// }

	// mag, err := strconv.ParseFloat(os.Args[1], 64)
	// if err == nil {
	// 	switch {
	// 	case mag < 2.0:
	// 		fmt.Println(mag, "is", "micro")
	// 	case mag == 2.0 || mag < 3.0:
	// 		fmt.Println(mag, "is", "very minor")
	// 	case mag == 3.0 || mag < 4.0:
	// 		fmt.Println(mag, "is", "minor")
	// 	case mag == 4.0 || mag < 5.0:
	// 		fmt.Println(mag, "is", "light")
	// 	case mag == 5.0 || mag < 6.0:
	// 		fmt.Println(mag, "is", "moderate")
	// 	case mag == 6.0 || mag < 7.0:
	// 		fmt.Println(mag, "is", "strong")
	// 	case mag == 7.0 || mag < 8.0:
	// 		fmt.Println(mag, "is", "major")
	// 	case mag == 8.0 || mag < 10.0:
	// 		fmt.Println(mag, "is", "great")
	// 	case mag <= 10.0:
	// 		fmt.Println(mag, "is", "massive")
	// 	}
	// }

	// const (
	// 	usage       = "Usage: [username] [password]"
	// 	errUser     = "Access denied for %q\n"
	// 	errPwd      = "Invalid password for %q\n"
	// 	accessOk    = "Access granted to %q\n"
	// 	user, user2 = "jack", "inanc"
	// 	pass, pass2 = "1888", "1879"
	// )

	// args := os.Args
	// if len(args) != 3 {
	// 	fmt.Println(usage)
	// 	return
	// }

	// u, p := os.Args[1], os.Args[2]

	// switch true {
	// case u != user && u != user2:
	// 	fmt.Printf(accessOk, u)
	// case (u == user && p == pass) || (u == user2 && p == pass2):
	// 	fmt.Printf(accessOk, u)
	// default:
	// 	fmt.Printf(errPwd, u)
	// }

	// args := os.Args
	// if len(args) != 3 {
	// 	fmt.Println("Avalable commands: lower, upper and title")
	// 	return
	// }

	// cmd, title := os.Args[1], os.Args[2]
	// switch cmd {
	// case "upper":
	// 	fmt.Printf("%q %scase is %q\n", title, cmd, strings.ToUpper(title))
	// case "lower":
	// 	fmt.Printf("%q %scase is %q\n", title, cmd, strings.ToLower(title))
	// default:
	// 	fmt.Println("You have choosen wrong command")
	// }

	y := time.Now().Year()
	leap := y%4 == 0 && (y%100 != 0 || y%400 == 0)
	days, month := 28, os.Args[1]

	switch m := strings.ToLower(month); m {
	case "april", "june", "september", "november":
		days = 30
	case "january", "march", "may", "july", "august", "october", "december":
		days = 31
	case "february":
		if leap {
			days = 29
		}
	default:
		fmt.Printf("%q is not a month\n", month)
	}

	fmt.Printf("%q has %d days.\n", month, days)
}
