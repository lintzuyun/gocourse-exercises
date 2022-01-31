package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Square struct {
	length float64
}

//attach a method to each that calculates AREA and returns it

func (r Circle) area() float64 {
	return math.Pi * r.radius * r.radius
}

func (s Square) area() float64 {
	return s.length * s.length
}

//create a type SHAPE that defines an interface as anything that has the AREA method

type shape interface {
	area() float64
}

//create a func INFO which takes type shape and then prints the area

func info(s shape) {
	x := s.area()
	fmt.Println(x)
}

func main() {

	// exercise #5
	circle1 := Circle{
		radius: 10,
	}

	square1 := Square{
		length: 16,
	}

	info(circle1)
	info(square1)
}
