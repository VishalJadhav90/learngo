package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	moodsArray := [...][3]string{
		{"What a nice day", "I am happy", "Wow, feels good"},
		{"I am sick", "Tired as hell", "Bad day"},
	}

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <name> <mood:(positive/negative)> ")
		return
	}

	name := os.Args[1]

	rand.Seed(time.Now().UnixNano())
	if os.Args[2] == "positive" {
		index := rand.Intn(3)
		fmt.Printf("%s %s\n", moodsArray[0][index], name)
	} else if os.Args[2] == "negative" {
		index := rand.Intn(3)
		fmt.Printf("%s %s\n", moodsArray[1][index], name)
	}
}
