package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(int(math.Pow(float64(a), 2)))
}
