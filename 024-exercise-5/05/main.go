package main

import "fmt"

func main() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(a)

	b := append(a[:3], a[6:]...)
	fmt.Println(b)
}
