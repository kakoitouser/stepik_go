package main

import (
	"fmt"
)

func main() {
	var text string
	fmt.Scan(&text)
	new := "0"
	for _, v := range text {
		if string(v) > new {
			new = string(v)
		}
	}
	fmt.Println(new)
}
