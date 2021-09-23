func vote(x int, y int, z int) int {
	zero, one := 0, 0
	arr := [3]int{x, y, z}
	for _, v := range arr {
		if v == 0 {
			zero++
		} else {
			one++
		}
	}
	if zero > one {
		return (0)
	} else {
		return (1)
	}
}