package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	fmt.Scan(&number)

	if number >= 100000 && number <= 999999 {
		number_str1 := strconv.Itoa(number / 1000)
		number_str2 := strconv.Itoa(number % 1000)

		number1 := 0
		number2 := 0
		for _, v := range number_str1 {
			num, _ := strconv.Atoi(string(v))
			number1 += num
		}
		for _, v := range number_str2 {
			num, _ := strconv.Atoi(string(v))
			number2 += num
		}
		if number1 == number2 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
