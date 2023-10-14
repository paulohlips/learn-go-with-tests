package arraysandslices

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(slicesToSum ...[]int) []int {
	var sums []int

	for _, slice := range slicesToSum {
		sums = append(sums, Sum(slice))
	}

	return sums
}

func SumAllTails(slicesToSumTail ...[]int) []int {
	var sums []int
	tail := []int{}
	for _, slice := range slicesToSumTail {
		if len(slice) != 0 {
			tail = slice[1:]
		}
		sums = append(sums, Sum(tail))
	}

	return sums
}