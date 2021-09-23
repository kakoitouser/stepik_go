package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)

	if number <= 10000 {
		if number%400 == 0 || (number%4 == 0 && number%100 != 0) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
