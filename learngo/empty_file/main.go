package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Provide a dictionary")
		return
	}
	fmt.Println("args...", args[1])
	files, err := ioutil.ReadDir(args[1])
	if err != nil {
		fmt.Println("Failed to fetch files ", files)
		return
	} else {
		total := 256 * len(files)
		names := make([]byte, 0, total)
		for _, file := range files {
			if !file.IsDir() && file.Size() == 0 {
				fmt.Printf("fileName: %s\n", file.Name())
				names = append(names, file.Name()...)
				names = append(names, '\n')
			}
		}
		fmt.Printf("names: %s\n", names)
		err := ioutil.WriteFile("/tmp/emptyfiles", names, 0644)
		if err != nil {
			fmt.Println("Failed to write file /tmp/emptyfiles due to error ", err)
		}
	}
}
