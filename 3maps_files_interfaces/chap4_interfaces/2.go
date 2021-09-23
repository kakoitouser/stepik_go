package main

type Battery string

func (b *Battery) String() string {
	count := 0
	for _, v := range *b {
		if string(v) == "1" {
			count++
		}
	}
	res := "["
	for i := 0; i < 10-count; i++ {
		res += " "
	}
	for i := 0; i < count; i++ {
		res += "X"
	}
	res += "]"
	return res
}

func main() {
	var test string = "test"
	batteryForTest := Battery(test)
}
