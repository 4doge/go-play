package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		remainder := i % 4
		fmt.Println(i, remainder)
	}
}
