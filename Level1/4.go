package main

import "fmt"

func main() {

	// exercise #4 -1
	type joy int

	// exercise #4 -2
	var x joy

	// exercise #4 -3
	fmt.Println(x)
	fmt.Printf("%T\n", x) // (1) Printf = format print (2) \n = line break

	x = 42
	fmt.Println(x)
}
