package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) calculateArea() float64 {
	return s.length * s.length
}

func (c circle) calculateArea() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	calculateArea() float64
}

func info(s shape) {
	fmt.Println(s.calculateArea())
}

func main() {
	c := circle{
		radius: 5,
	}

	s := square{
		length: 10,
	}

	info(c)
	info(s)
}
