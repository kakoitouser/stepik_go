package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var a, b float64
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		errors.New("fuck u")
	} else {
		res, err := FindTheHypotenuze(a, b)
		if err != nil {
			fmt.Println("hey yo")
		} else {
			fmt.Println(res)
		}
	}
}
func FindTheHypotenuze(a, b float64) (float64, error) {
	if a < 0 || b < 0 {
		return -1, errors.New("a or b les than 0")
	} else {
		return math.Sqrt(a*a + b*b), nil
	}
}
