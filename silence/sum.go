package silence

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return sum

}
