package main

import "fmt"

func main() {
	bob := []string{"Bob", "IT guy"}
	alice := []string{"Alice", "Instagirl"}
	fmt.Println(bob)
	fmt.Println(alice)
	coop := [][]string{bob, alice}
	fmt.Println(coop)
}
