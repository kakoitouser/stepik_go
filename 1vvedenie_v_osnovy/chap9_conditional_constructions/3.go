package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)

	if number <= 10000 {
		for number >= 10 {
			number /= 10
		}
		fmt.Println(number)
	}
}
