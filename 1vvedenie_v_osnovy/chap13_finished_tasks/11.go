package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n > 0 && n < 100 {
		if (n >= 5 && n <= 20) || (n%10 == 0) || (n%10 >= 5 && n%10 <= 9) {
			fmt.Println(n, "korov")
		} else if n%10 == 1 {
			fmt.Println(n, "korova")
		} else if n%10 >= 2 && n%10 <= 4 {
			fmt.Println(n, "korovy")
		}
	}
}
