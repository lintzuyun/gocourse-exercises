package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// A method is a function with a special receiver argument

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "and I am ", p.age, "years old.")
}

func main() {

	// exercise #4
	p1 := person{
		first: "Joy",
		last:  "Lin",
		age:   30}

	p1.speak() // call the method from the value of type person
}
