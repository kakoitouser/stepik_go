package main

import (
	"fmt"
)

func main() {
	var num, n, sum int
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		fmt.Scan(&n)
		if n >= 10 && n <= 99 && n%8 == 0 {
			sum += n
		}
	}
	fmt.Println(sum)
}
