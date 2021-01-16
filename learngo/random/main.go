package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	var guess = 10

	t := time.Now()

	if len(os.Args) != 2 {
		fmt.Println("Please enter a random number as argument")
		return
	}

	maxTurns, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Please enter a valid number")
		return
	}

	rand.Seed(t.UnixNano())
	for try := 0; try < maxTurns; try++ {
		val := rand.Intn(guess + 1)
		if val == guess {
			if try == 0 {
				fmt.Println("You won in fist attempt !! thats great")
				return
			}
			msgNum := rand.Intn(3)
			switch msgNum {
			case 0:
				fmt.Println("You are taking prize with you, yay ^.^")
			case 1:
				fmt.Println("You're winner $")
			case 2:
				fmt.Println("You've got it man!")
			default:
				fmt.Println("Thats a win for you..")
			}
			return
		}
	}
	msgNum := rand.Intn(3)
	switch msgNum {
	case 0:
		fmt.Println("Hard luck.")
	case 1:
		fmt.Println("Next time try hard")
	case 2:
		fmt.Println("You loose")
	default:
		fmt.Println("You can go home")
	}
}
