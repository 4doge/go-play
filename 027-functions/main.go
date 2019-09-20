package main

import "fmt"

func main() {
	greeting, randomNum := sayMyName("Heisenberg")
	fmt.Println(greeting, randomNum)
	printNums(1, 2, 3, 4, 5)
}

func sayMyName(name string) (string, int) {
	return fmt.Sprint("Hey, you're name is ", name), 13
}

func printNums(nums ...int) {
	fmt.Println(nums)
	fmt.Printf("%T\n", nums)
}