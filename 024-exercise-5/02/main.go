package main

import "fmt"

func main() {
	a := []int{0, 11, 22, 33, 44, 55, 66, 77, 88, 99}
	fmt.Println(a)

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", a)
}
