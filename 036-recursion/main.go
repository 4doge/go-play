package main

import "fmt"

func main() {
	f := factorial(5)
	fmt.Println(f)
	f2 := factorial2(5)
	fmt.Println(f2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorial2(n int) int {
	total := 1
	for i := 1; i <= n; i++ {
		total *= i
	}
	return total
}