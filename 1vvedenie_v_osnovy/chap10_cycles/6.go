package main

import (
	"fmt"
)

func main() {
	var n int
	// считываем числа пока не будет введен 0
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		if n > 100 {
			break
		} else if n >= 10 {
			fmt.Println(n)
		} else {
			continue
		}
	}
}
