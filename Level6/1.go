package main

import "fmt"

func main() {

	// exercise #1

	// fmt.Println(foo())
	// fmt.Println(bar())

	x := foo()
	y, z := bar()
	fmt.Println(x, y, z)
}

func foo() int {
	return 40
}

func bar() (int, string) {
	return 40, "Hello"
}
