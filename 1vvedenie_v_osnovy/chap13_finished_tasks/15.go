package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, del int
	fmt.Scan(&n, &del)

	arr := make([]int, 0)
	if n > 0 {
		str := strconv.Itoa(n)
		for i := range str {
			delStr := strconv.Itoa(del)
			if string(str[i]) == delStr {
				arr = append(arr, i)
			}
		}
		ind := 0
		for _, v := range arr {
			v -= ind
			ind++
			str = str[:v] + str[v+1:]
		}
		num, _ := strconv.Atoi(string(str))
		fmt.Println(num)
	}
}
