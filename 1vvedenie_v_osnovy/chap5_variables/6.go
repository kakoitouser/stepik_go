package main

import (
	"fmt"
)

func main() {
	var d int
	fmt.Scan(&d)
	if d > 0 && d < 360 {
		h := d / 30
		m := (d - h*30) * 2
		fmt.Println("It is", h, "hours", m, "minutes.")
	}
}
