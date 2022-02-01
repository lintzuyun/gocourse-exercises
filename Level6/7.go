package main

import "fmt"

func main() {

	// exercise #7
	x := func() {
		fmt.Println("Hello world!")
	}

	x()

	fmt.Printf("%T\n", x)
}
