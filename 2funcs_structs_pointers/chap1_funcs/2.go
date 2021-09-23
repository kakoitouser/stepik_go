func minimumFromFour() int {
	var arr [4]int
	min := 9999999999
	for ind := range arr {
		fmt.Scan(&arr[ind])
		if min > arr[ind] {
			min = arr[ind]
		}
	}
	return min
}