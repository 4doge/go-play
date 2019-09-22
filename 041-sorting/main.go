package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	occupation string
	age int
}

type byAge []person

func (x byAge) Len() int {
	return len(x)
}

func (x byAge) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x byAge) Less(i, j int) bool {
	return x[i].age < x[j].age
}

type byName []person

func (x byName) Len() int {
	return len(x)
}

func (x byName) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x byName) Less(i, j int) bool {
	return x[i].name < x[j].name
}

func main() {
	nums := []int{111, 222, 11, 22, 33, 44, 55, 66, 77, 88, 99}
	strs := []string{"Yayee", "Alice", "Bob", "Eve", "John"}

	fmt.Println("Original:", nums)
	fmt.Println("Original:", strs)

	sort.Ints(nums)
	fmt.Println("Sorted ASC:", nums)
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println("Sorted DESC:", nums)
	

	sort.Strings(strs)
	fmt.Println("Sorted ASC:", strs)
	sort.Sort(sort.Reverse(sort.StringSlice(strs)))
	fmt.Println("Sorted DESC:", strs)

	p := person{
		name: "Bob",
		occupation: "IT guy",
		age: 27,
	}

	p2 := person{
		name: "Alice",
		occupation: "Instagirl",
		age: 24,
	}

	p3 := person{
		name: "Eve",
		occupation: "Listener",
		age: 23,
	}

	p4 := person{
		name: "John",
		occupation: "Regular guy",
		age: 44,
	}

	people := []person{p, p2, p3, p4}
	fmt.Println("Original:", people)

	sort.Sort(byAge(people))
	fmt.Println("Sorted by age, ASC:", people)

	sort.Sort(byName(people))
	fmt.Println("Sorted by name, ASC:", people)
}
