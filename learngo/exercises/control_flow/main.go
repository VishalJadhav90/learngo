package main

import (
	"fmt"
	"os"
)

func main() {
	// age, _ := strconv.ParseInt(os.Args[1], 10, 64)
	// if age > 60 {
	// 	fmt.Println("Getting Older")
	// } else if age > 30 {
	// 	fmt.Println("Getting Wiser")
	// } else if age > 20 {
	// 	fmt.Println("Adulthood")
	// } else {
	// 	fmt.Println("Booting up")
	// }

	// isSphere, radius := true, 200
	// big := false
	// if radius > 200 {
	// 	big = true
	// }

	// if isSphere && big {
	// 	fmt.Println("Its a big sphere")
	// } else {
	// 	fmt.Println("I dont know")
	// }

	username, password := "jack", "1888"

	args := len(os.Args)
	if args < 3 {
		fmt.Println("Usage: [username] [password]")
	} else {
		inUsername, inPassword := os.Args[1], os.Args[2]
		if inUsername != username {
			fmt.Println("Access denied for", inUsername)
		} else if inUsername == username && inPassword != password {
			fmt.Println("Invalid password for", inUsername)
		} else {
			fmt.Println("Access granted to", inUsername)
		}
	}

}
