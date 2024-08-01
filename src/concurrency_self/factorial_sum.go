package concurrency

func factorial(n int, data chan<- int) {
	product := 1
	for i := range n + 1 {
		if i == 0 {
			continue
		}
		product *= i
	}
	data <- product
}

func FactorialSum(n int) int {
	factorialChan := make(chan int)

	for i := range n + 1 {
		if i == 0 {
			continue
		}
		go factorial(i, factorialChan)
	}

	sum := 0
	for value := range factorialChan {
		sum += value
	}
	return sum
}
