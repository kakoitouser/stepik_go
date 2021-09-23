package main

import (
	"fmt"
)

func main() {
	var text string
	fmt.Scan(&text)
	fmt.Println(NewStr(text))
}
func NewStr(text string) string {
	newstr := ""
	for i := 0; i < len(text); i++ {
		newstr += string(text[i]) + "*"
	}
	newstr = newstr[:len(newstr)-1]
	return newstr
}
