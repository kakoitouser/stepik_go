package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	max := -9999999999
	for _, v := range array {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}
