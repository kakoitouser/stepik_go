package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	if N >= 4 {
		arr := make([]int, N, N)
		for ind := range arr {
			fmt.Scan(&arr[ind])
		}
		fmt.Println(arr[3])
	}
}
