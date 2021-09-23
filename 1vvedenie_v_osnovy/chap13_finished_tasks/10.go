package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scan(&a, &b)
	count := 0
	for i := b; i >= a; i-- {
		if i%7 == 0 {
			count = 1
			fmt.Println(i)
			break
		}
	}
	if count == 0 {
		fmt.Println("NO")
	}
}
