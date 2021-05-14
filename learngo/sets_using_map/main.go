package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	set := map[string]bool{}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		if word == "quit" {
			break
		}
		set[word] = true
	}
	fmt.Println("enter word to check")
	scanner.Scan()
	wordToCheck := scanner.Text()
	if _, exists := set[wordToCheck]; !exists {
		fmt.Printf("%s does not exists\n", wordToCheck)
	} else {
		fmt.Printf("%s exists\n", wordToCheck)
	}
}
