package main

import "fmt"

func CalculateFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * CalculateFactorial(n-1)
}

func main() {
	fmt.Println(CalculateFactorial(5)) // Output: 120
}
