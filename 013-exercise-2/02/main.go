package main

import "fmt"

func main() {
	a := ("Hello" == "Hello")
	b := (10 <= 13)
	c := (27 >= 4)
	d := ("Hello" != "Hey")
	e := (14 < 17)
	f := (31 > 9)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
