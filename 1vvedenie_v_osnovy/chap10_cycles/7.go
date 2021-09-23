package main

import (
	"fmt"
)

func main() {
	var x, p, y, year int
	fmt.Scan(&x, &p, &y)
	x *= 100
	y *= 100

	for x <= y {
		x = x + (x * p / 100)
		year++
	}
	fmt.Println(year)
}
