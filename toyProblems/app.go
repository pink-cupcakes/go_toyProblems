package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

func main() {
	fibVal := make(map[int]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		fibVal[i] = fibonacci(i, fibVal)
	}
}

func fibonacci(i int, fibVal map[int]int) int {
	if fibVal[i] != 0 {
		fmt.Printf("Catch case seen i: %d and value is: %d\n", i, fibVal[i])
		return fibVal[i]
	}
	if i < 1 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonacci(i - 1, fibVal) + fibonacci(i - 2, fibVal)
}