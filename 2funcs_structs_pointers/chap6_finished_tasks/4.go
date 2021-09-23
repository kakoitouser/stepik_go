package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		text string
		num  int
	)

	fmt.Scan(&text)
	new := ""
	for _, v := range text {
		num, _ = strconv.Atoi(string(v))
		new += strconv.Itoa(num * num)
	}
	fmt.Println(new)
}
