package main

import "fmt"

func main() {
	favSport := "coding"
	switch favSport {
	case "football":
		fmt.Println("Football")
	case "hockey":
		fmt.Println("Hockey")
	default:
		fmt.Println("You don't have favorite sport")
	}
}
