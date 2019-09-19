package main

import "fmt"

func main() {
	// composite literal
	// x := type{values}
	// slice = group values of the same type
	x := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n%T\n", x, x)
	fmt.Printf("Length: %d\n", len(x))
	fmt.Println(x[2])

	for index, value := range x {
		fmt.Println(index, value)
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	// [ include : notInclude ]
	fmt.Println(x)
	// fmt.Println(x[1:])
	// fmt.Println(x[1:3])

	x = append(x, 13, 14, 15, 16)
	fmt.Println(x)

	y := []int{123, 456, 789}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:4], x[6:]...)
	fmt.Println(x)


}
