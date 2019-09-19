package main

import "fmt"

func main() {
	m := map[string][]string{
		"bob":   []string{"it", "women", "alice"},
		"alice": []string{"insta", "men", "bob"},
		"eve":   []string{"nothing"},
	}
	fmt.Println(m)
	delete(m, "eve")
	fmt.Println(m)
	m["carol"] = []string{"crypto", "men"}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
