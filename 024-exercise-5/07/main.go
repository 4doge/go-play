package main

import "fmt"

func main() {
	people := [][]string{
		[]string{"Bob", "Smith", "IT guy"},
		[]string{"Alice", "Smith", "Instagirl"},
	}
	fmt.Println(people)
	for _, v := range people {
		fmt.Println(v)
		for _, p := range v {
			fmt.Println(p)
		}
	}
}
