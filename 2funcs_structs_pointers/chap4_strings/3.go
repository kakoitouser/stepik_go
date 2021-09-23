package main

import (
	"fmt"

	"strings"
)

func main() {

	x, s := "", ""

	fmt.Scan(&x)

	fmt.Scan(&s)

	res := strings.Index(x, s)

	fmt.Println(res)

}
