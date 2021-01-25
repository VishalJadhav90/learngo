package main

import "fmt"

func main() {
	wizards := [...][3]string{
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	fmt.Printf("%-15s%-15s%-15s\n", "Firstname", "Lastname", "Nickname")
	fmt.Printf("====================================================\n")
	for i := range wizards {
		for j := range wizards[i] {
			fmt.Printf("%-15s", wizards[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("----------------------------------------------------\n")
}
