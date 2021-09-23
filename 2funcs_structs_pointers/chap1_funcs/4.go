func fibonacci(n int) int {
	cur := 1
	if n > 2 {
		return fibonacci(n-1) + fibonacci(n-2)
	}
	return cur
}