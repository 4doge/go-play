package main

import "fmt"

const now = 2019

const (
	_ = iota
	a = now + iota
	b = now + iota
	c = now + iota
	d = now + iota
)

func main() {
	fmt.Println(now)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
