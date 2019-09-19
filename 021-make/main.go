package main

import "fmt"

func main() {
	x := make([]int, 5, 7)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 13)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 14)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 15)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
