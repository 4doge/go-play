package main

import "fmt"

const (
	_ = iota
	kbi = 1 << (iota * 10)
	mbi = 1 << (iota * 10)
	gbi = 1 << (iota * 10)
)

func main() {
	x := 4
	fmt.Printf("%v %T %b\n", x, x, x)
	y := x << 1
	fmt.Printf("%v %T %b\n", y, y, y)

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb
	fmt.Printf("%v %b\n", kb, kb)
	fmt.Printf("%v %b\n", mb, mb)
	fmt.Printf("%v %b\n", gb, gb)

	fmt.Printf("%v %b\n", kbi, kbi)
	fmt.Printf("%v %b\n", mbi, mbi)
	fmt.Printf("%v %b\n", gbi, gbi)
}
