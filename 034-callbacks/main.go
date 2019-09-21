package main

import "fmt"

func main() {
	nums := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	s := sum(nums...)
	fmt.Println(s)
	s2 := sumEven(sum, nums...)
	fmt.Println(s2)
	s3 := sumOdd(sum, nums...)
	fmt.Println(s3)
}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func sumEven(f func(nums ...int) int, nums...int) int {
	evenNums := []int{}
	for _, v := range nums {
		if v % 2 == 0 {
			evenNums = append(evenNums, v)
		}
	}
	return f(evenNums...)
}

func sumOdd(f func(nums ...int) int, nums...int) int {
	oddNums := []int{}
	for _, v := range nums {
		if v % 2 != 0 {
			oddNums = append(oddNums, v)
		}
	}
	return f(oddNums...)
}