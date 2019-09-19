package main

import "fmt"

func main() {
	year := 1980
	now := 2019
	for year <= now {
		fmt.Println(year)
		year++
	}
}
