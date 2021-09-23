func adding(a, b string) int64 {
	var aint, bint int
	var astr, bstr string
	for _, v := range a {
		if unicode.IsDigit(v) {
			astr += string(v)
		}
	}
	for _, v := range b {
		if unicode.IsDigit(v) {
			bstr += string(v)
		}
	}
	aint, _ = strconv.Atoi(astr)
	bint, _ = strconv.Atoi(bstr)
	return int64(aint) + int64(bint)
}