package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number uint
	fmt.Scan(&number)

	fn := func(number uint) uint {
		result := ""
		strNum := strconv.Itoa(int(number))
		for _, v := range strNum {
			val, err := strconv.Atoi(string(v))
			if err != nil {
				panic("panic")
			}
			if val%2 == 0 && val != 0 {
				result += strconv.Itoa(val)
			}
		}
		res, _ := strconv.Atoi(result)
		if res == 0 {
			return uint(100)
		} else {
			return uint(res)
		}
	}
	fn(number)
}
