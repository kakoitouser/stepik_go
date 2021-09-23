package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)

	if number >= 100 && number <= 999 {
		h := number / 100
		t := number / 10 % 10
		o := number % 10
		if h != t && t != o && h != o {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
