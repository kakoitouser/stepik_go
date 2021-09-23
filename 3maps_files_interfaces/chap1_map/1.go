var nums [10]int
var num int
mapa := make(map[int]int)
for i, _ := range nums {
	fmt.Scan(&num)
	if val, ok := mapa[num]; ok {
		nums[i] = val
	} else {
		nums[i] = work(num)
		mapa[num] = nums[i]
	}
}
for _, v := range nums {
	fmt.Print(v, " ")
}