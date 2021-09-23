package main

import (
	"fmt"
	"strings"
)

func newstr(text string) string {
	newtext := ""
	for _, v := range text {
		if strings.Count(text, string(v)) < 2 {
			newtext += string(v)
		}
	}
	return newtext
}

func main() {
	var text string
	fmt.Scan(&text)
	fmt.Println(newstr((text)))
}
