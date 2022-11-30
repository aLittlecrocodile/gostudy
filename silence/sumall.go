package silence

func SumAll(numberToSum ...[]int) (sums []int) {
	lengthOfNumber := len(numberToSum)
	sums = make([]int, lengthOfNumber)
	for i, numbers := range numberToSum {
		sums[i] = Sum(numbers)

	}
	return

}
