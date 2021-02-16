package main

import (
	"bytes"
	"fmt"
)

func main() {
	png, header := []byte{'P', 'N', 'G'}, []byte{}
	header = append(header, png...)
	result := bytes.Compare(png, header)
	fmt.Printf("%#v %#v are they equal... %t\n", header, png, result == 0)
}
