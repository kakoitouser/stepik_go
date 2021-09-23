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

		quantity := 0
		for _, v := range arr {
			if v > 0 {
				quantity++
			}
		}
		fmt.Println(quantity)
	}
}
