package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	number := 0
	if a < 100 && b < 100 && a < b {
		for a <= b {
			number += a
			a++
		}
		fmt.Println(number)
	}
}
