func sumInt(args ...int) (numbersOfArguments int, sum int) {
	numbersOfArguments = len(args)
	if numbersOfArguments > 0 {
		for _, value := range args {
			sum += value
		}
		return numbersOfArguments, sum
	}
	return 0, 0
}