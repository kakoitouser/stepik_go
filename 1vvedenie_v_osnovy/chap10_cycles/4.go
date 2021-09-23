package main

import (
	"fmt"
)

func main() {
	var n int
	max := 0
	quant := 0
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		if n > max {
			max = n
			quant = 1
		} else if n == max {
			quant++
		}
	}
	fmt.Println(quant)
}
