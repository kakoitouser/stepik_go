package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)

	if 1 <= k && k <= 86398 {
		h := k / 3600
		m := (k - h*3600) / 60

		fmt.Printf("It is %v hours %v minutes.", h, m)
	}
}
