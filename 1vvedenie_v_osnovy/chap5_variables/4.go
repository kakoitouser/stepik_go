package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	if a < 10000 {
		fmt.Println(a % 10)
	}
}
