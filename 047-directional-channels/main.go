package main

func main() {
	c := make(chan<- int, 2)
	c <- 13
	c <- 99
	// can't read from `send-only` channel
	// fmt.Println(<-c)

	// c2 := make(<-chan int, 2)
	// can't send to `receive-only` channel
	// c2 <- 13
	// c2 <- 13
	// fmt.Println(<-c2)
}
