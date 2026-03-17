package main

func Sum(numbers []int) any {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
