package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumTail(numbers []int) (tailSum int) {
	if len(numbers) <= 1 {
		return 0
	}
	tail := numbers[1:]
	return Sum(tail)
}

func SumAllTails(intSlices ...[]int) (tailSums []int) {
	for _, numbers := range intSlices {
		tailSums = append(tailSums, SumTail(numbers))
	}
	return tailSums
}