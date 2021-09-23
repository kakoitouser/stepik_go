package main

import "fmt"

func main() {
	// package main уже объявлен, нужные импорты уже есть
	var a, b int
	fmt.Scan(&a, &b)
	res, err := divide(a, b)
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println("ошибка")
	}
}

func divide(a, b int) (int, error) {
	return a + b, nil
}
