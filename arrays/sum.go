package main

func SumSlice(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumArray(numbers [5]int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	return nil
}
