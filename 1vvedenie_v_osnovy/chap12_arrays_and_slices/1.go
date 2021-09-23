var workArray [10]uint8
var arr [6]uint8

for ind := range workArray {
	fmt.Scan(&workArray[ind])
}
for ind := range arr {
	fmt.Scan(&arr[ind])
}

for i := 0; i < 6; i += 2 {
	temp := workArray[arr[i]]
	workArray[arr[i]] = workArray[arr[i+1]]
	workArray[arr[i+1]] = temp
}

for _, v := range workArray {
	fmt.Print(v, " ")
}