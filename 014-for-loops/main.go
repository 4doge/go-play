package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("i: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\tj: %d\n", j)
		}
	}

	a := 1
	for a <= 10 {
		fmt.Println("a", a)
		a++
	}

	b := 1
	for {
		b++
		if b > 100 {
			break
		}
		if b % 2 != 0 {
			continue
		}
		fmt.Println("b", b)
	}
}
