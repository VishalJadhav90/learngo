package main

import (
	"fmt"
	"os"
	"strings"
)

const protocol = "https://"

func main() {
	urlStr := os.Args[1]
	byteArr := make([]byte, len(strings.ToLower(urlStr)))
	copy(byteArr, strings.ToLower(urlStr))
	//fmt.Println("len(protocol)", len(protocol))
	var found bool = false
	for i := 0; i < len(byteArr); i++ {
		//fmt.Printf("blah is this %s\n", string(byteArr[i:len(protocol)]))
		if i+len(protocol) < len(urlStr) && string(byteArr[i:i+len(protocol)]) == protocol {
			fmt.Println("we have found this....")
			found = true
			i = i + len(protocol) - 1
			continue
		}
		if found {
			if byteArr[i] == ' ' {
				found = false
				continue
			}
			byteArr[i] = '*'
		}
	}
	fmt.Printf("masked string...%s\n", byteArr)
}
