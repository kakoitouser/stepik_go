package main

import (
	"fmt"
	"unicode/utf8"
)

func isPalindrom(text string) string {
	var reverseText string
	var reverseRune []rune

	runes := []rune(text)
	for i := utf8.RuneCountInString(text) - 1; i >= 0; i-- {
		reverseRune = append(reverseRune, runes[i])
	}
	reverseText = string(reverseRune)

	if text == reverseText {
		return "Палиндром"
	} else {
		return "Нет"
	}
}

func main() {
	var text string
	fmt.Scan(&text)
	fmt.Println(isPalindrom((text)))
}
