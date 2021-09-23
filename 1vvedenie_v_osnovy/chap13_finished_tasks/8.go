package main

import "fmt"

func main() {
	var n, num, qu int
	min := 99999999999
	fmt.Scan(&n)

	if n > 0 {
		for i := 0; i < n; i++ {
			fmt.Scan(&num)
			if num < min {
				min = num
				qu = 0
			}
			if num == min {
				qu++
			}
		}
		fmt.Println(qu)
	}
}
