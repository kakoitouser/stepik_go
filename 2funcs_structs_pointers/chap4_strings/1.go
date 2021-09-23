package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runes := []rune(text)
	if unicode.IsUpper(runes[0]) {
		if runes[utf8.RuneCountInString(text)-1] == '.' {
			fmt.Println("Right")
			return
		}
		fmt.Println("Wrong")
		return
	}
	fmt.Println("Wrong")
}
