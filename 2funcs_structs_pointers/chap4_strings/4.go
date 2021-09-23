package main

import (
	"fmt"
	"unicode/utf8"
)

func newstr(text string) string {
	var newRune []rune

	runes := []rune(text)
	for i := 1; i < utf8.RuneCountInString(text); i += 2 {
		newRune = append(newRune, runes[i])
	}
	return string(newRune)
}

func main() {
	var text string
	fmt.Scan(&text)
	fmt.Println(newstr((text)))
}
