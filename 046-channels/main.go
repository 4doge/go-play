package main

import "fmt"

func main() {
	c := make(chan int, 2)
	// go func() {
	// 	c <- 13
	// }()
	c <- 13
	c <- 99
	fmt.Println(<-c)
	fmt.Println(<-c)
}
