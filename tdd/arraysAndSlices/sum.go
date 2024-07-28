package arraysandslices

func Sum(numbers []int) int {
	// sum := 0
	// for i := 0; i < len(numbers); i++ {
	// 	sum += numbers[i]
	// }
	// return sum

	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

// func SumAll(numbersToSum ...[]int) []int {
// 	lengthOfNumbers := len(numbersToSum)
// 	sums := make([]int, lengthOfNumbers)

// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}
// 	return sums
// }

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	// var sums []int
	// for _, numbers := range numbersToSum {
	// 	if len(numbers) == 0 {
	// 		sums = append(sums, 0)
	// 	} else {
	// 		tail := numbers[1:]
	// 		sums = append(sums, Sum(tail))
	// 	}
	// }
	// return sums

	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(numbersToSum, sumTail, []int{})
}
