package main

import (
	"fmt"
)

func main() {
	var A int
	fmt.Scan(&A)
	a := 0
	b := 1
	i := 1
	temp := 0
	if A > 1 {
		for {
			temp = a + b
			a = b
			b = temp
			i++
			if b == A {
				fmt.Println(i)
				break
			}
			if b > A {
				fmt.Println(-1)
				break
			}
		}
	}
}
