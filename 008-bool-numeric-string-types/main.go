package main

import "fmt"

func main() {
	// Boolean
	a := 3
	b := 13
	fmt.Println(a != b)

	// Numbers
	x := 13
	y := 13.1234

	fmt.Printf("%v %v %T %T\n", x, y, x, y)

	// String
	s := "Hello world string!"
	ms := `Hello world 
	multiline
	string`
	fmt.Printf("%v %T\n", s, s)
	fmt.Printf("%v %T\n", ms, ms)

	bs := []byte(s)
	fmt.Printf("%v %T\n", bs, bs)

	for  i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}

	for i, v := range s {
		fmt.Println(i, v)
		fmt.Printf("Index: %d | Hex value: %#x \n", i, v)
	}

}
