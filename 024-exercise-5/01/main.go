package main

import "fmt"

func main() {
	a := [5]int{}
	fmt.Println(a)

	for i := range a {
		a[i] = i
	}

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", a)

}
