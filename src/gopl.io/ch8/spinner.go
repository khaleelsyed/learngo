package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n int = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci (%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, char := range `-\|/` {
			fmt.Printf("\r%c", char)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
