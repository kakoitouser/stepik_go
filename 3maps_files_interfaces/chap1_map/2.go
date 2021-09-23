
repetitive := make([]string, 0)
for k := range cityPopulation {
	for _, v := range groupCity[10] {
		if k == v {
			repetitive = append(repetitive, v)
		}
	}
	for _, v := range groupCity[1000] {
		if k == v {
			repetitive = append(repetitive, v)
		}
	}
}

for _, val := range repetitive {
	delete(cityPopulation, val)
}