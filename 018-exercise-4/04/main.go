package main

import "fmt"

func main() {
	year := 1980
	now := 2019
	for {
		if year <= now {
			fmt.Println(year)
			year++
		} else {
			break
		}
	}
}
