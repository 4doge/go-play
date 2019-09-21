package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	f := foo(nums...)
	fmt.Println(f)
	b := bar(nums)
	fmt.Println(b)
}

func foo(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func bar(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
