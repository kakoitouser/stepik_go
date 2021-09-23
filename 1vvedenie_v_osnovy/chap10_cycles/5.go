package main

import (
	"fmt"
)

func main() {
	var n, c, d int
	fmt.Scan(&n, &c, &d)

	if n < 10000 && c < 10000 && d < 10000 && n > 0 && c > 0 && d > 0 {
		for i := 1; i <= n; i++ {
			if i%c == 0 && i%d != 0 {
				fmt.Println(i)
				break
			}
		}
	}
}
