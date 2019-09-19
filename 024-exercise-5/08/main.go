package main

import "fmt"

func main() {
	m := map[string][]string{
		"bob":   []string{"it", "women", "alice"},
		"alice": []string{"insta", "men", "bob"},
		"eve":   []string{"nothing"},
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
}
