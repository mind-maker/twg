package math

// Sum sum all numbers passed in the argument slice
func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}

	return sum
}

// Add add two numbers
func Add(a, b int) int {
	return a + b
}
