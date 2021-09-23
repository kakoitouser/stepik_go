package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	if 1 <= num && num <= 100 {

		arr := make([]int, num)
		for i := range arr {
			fmt.Scan(&arr[i])
		}

		for i := 0; i < len(arr); i += 2 {
			fmt.Print(arr[i], " ")
		}
	}
}
