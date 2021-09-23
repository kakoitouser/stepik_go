package main

import (
	"fmt"
	"unicode"
)

func newstr(text string) string {
	if len(text) >= 5 {
		for _, v := range text {
			if unicode.Is(unicode.Latin, v) == false && unicode.IsDigit(v) == false {
				return "Wrong password"
			}
		}
		return "Ok"
	} else {
		return "Wrong password"
	}
}

func main() {
	var text string
	fmt.Scan(&text)
	fmt.Println(newstr((text)))
}
