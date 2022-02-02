package main

import "fmt"

type person struct {
	name string
}

func main() {
	p1 := person{"Joy"}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

// create a func called “changeMe” with a *person as a parameter
// change the value stored at the *person address

func changeMe(p *person) {
	p.name = "Craig"
}
