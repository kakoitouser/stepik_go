package main

import (
	"fmt"
	"math"
)

func main() {
	var n, sum int
	fmt.Scan(&n)

	if n < int(math.Pow(10, 7)) {
		for {
			if n == 0 && sum < 10 {
				fmt.Println(sum)
				break
			} else if n == 0 && sum >= 10 {
				n = sum
				sum = 0
			}
			l := n % 10
			n /= 10
			sum += l
		}
	}
}
