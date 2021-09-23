package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)
	var arr []int

	if 0 <= num1 && num1 <= 10000 && 0 <= num2 && num2 <= 10000 {
		n1str := strconv.Itoa(num1)
		n2str := strconv.Itoa(num2)

		for _, v := range n1str {
			for _, v2 := range n2str {
				if v == v2 {
					n, _ := strconv.Atoi(string(v))
					arr = append(arr, n)
				}
			}
		}
		for _, v := range arr {
			fmt.Print(v, " ")
		}
	}
}
