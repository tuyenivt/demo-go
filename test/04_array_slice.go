package main

func Sum(numbers [5]int) any {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
