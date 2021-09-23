package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)
	number := 0
	for i := 0; ; i++ {
		number = int(math.Pow(2, float64(i)))
		if number > N {
			break
		}
		fmt.Print(number, " ")
	}
}
