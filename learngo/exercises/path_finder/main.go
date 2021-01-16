package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	val := os.Environ()
	var pathEnv string
	for _, field := range val {
		//fmt.Println(field)
		if strings.HasPrefix(field, "PATH=") {
			fmt.Println(field)
			pathEnv = strings.Split(field, "=")[1]
		}
	}
	//fmt.Println("PATH:", pathEnv)
	for i, path := range strings.Split(pathEnv, ":") {
		if strings.Contains(path, os.Args[1]) {
			fmt.Printf("#%d %s\n", i+1, path)
		}
	}
}
