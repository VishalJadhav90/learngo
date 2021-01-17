package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var moods = [...]string{
		"what a happy day",
		"I feel terrible",
		"Sick as hell",
		"Yay Just had s..",
		"Busy day",
	}

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(5)
	fmt.Printf("\n%v %v\n", moods[i], os.Args[1])
}
