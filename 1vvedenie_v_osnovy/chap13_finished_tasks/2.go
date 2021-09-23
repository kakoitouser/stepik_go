package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	if 100 <= num && num <= 999 {
		h := num / 100
		t := num % 100 / 10
		l := num % 10

		fmt.Printf("%d%d%d", l, t, h)
	}
}
