package main

import "fmt"

func main() {

	// exercise #4

	var a int = 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

	var b = (a << 1) // shifts the bits of "a" over 1 position to the left and assign it to "b"
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)
}
