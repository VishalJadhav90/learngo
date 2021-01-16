package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const maxTrials = 10

const usage = `You should enter command as
go run main.go -v num1 num2`

func main() {

	rand.Seed(time.Now().UnixNano())

	if len(os.Args) < 4 {
		fmt.Println(usage)
		return
	}

	if os.Args[1] != "-v" {
		fmt.Println(usage)
		return
	}

	firstGuess, err1 := strconv.Atoi(os.Args[len(os.Args)-2])
	secondGuess, err2 := strconv.Atoi(os.Args[len(os.Args)-1])
	if err1 != nil || err2 != nil {
		fmt.Println("Please enter valid numbers as guesses")
		return
	}

	max := secondGuess
	if firstGuess > secondGuess {
		max = firstGuess
	}

	for trial := 0; trial < max; trial++ {
		n := rand.Intn(max + 1)
		fmt.Printf("%d ", n)
		if n == firstGuess {
			fmt.Printf("You've guessed first number %d\n", firstGuess)
			return
		} else if n == secondGuess {
			fmt.Printf("You've guessed second number %d\n", secondGuess)
			return
		}
	}
	fmt.Printf("Man, You couldnt guess either %d or %d\n", firstGuess, secondGuess)
}
