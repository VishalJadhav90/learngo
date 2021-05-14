package main

import "fmt"

func main() {
	dict := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mukemmel",
	}

	tdict := map[string]string{}

	for k, v := range dict {
		tdict[v] = k
	}
	if value, ok := tdict["iyi"]; !ok {
		fmt.Println("hi key not found in dict")
		return
	} else {
		fmt.Println("iyi in turkish is", value)
	}

	delete(tdict, "iyi")
	fmt.Println("tdict...", tdict)

}
