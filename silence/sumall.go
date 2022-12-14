package silence

func SumAll(numberToSum ...[]int) []int {
	var sums []int
	//lengthOfNumber := len(numberToSum)
	sums = make([]int, 0)
	for _, numbers := range numberToSum {
		sums = append(sums, Sum(numbers))

	}
	return sums

}
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)

		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}

	}

	return sums
}
