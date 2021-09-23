package main

import "fmt"

func main() {
	var n, num, qu int
	fmt.Scan(&n)

	if n > 0 {
		for i := 0; i < n; i++ {
			fmt.Scan(&num)
			if num == 0 {
				qu++
			}
		}
		fmt.Println(qu)
	}
}
