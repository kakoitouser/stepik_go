func test(x1 *int, x2 *int) {
	var temp int = *x1
	*x1 = *x2
	*x2 = temp
	fmt.Println(*x1, *x2)
}