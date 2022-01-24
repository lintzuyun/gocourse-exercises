package main

import (
	"fmt"
)

func main() {

	// exercise #2
	// using a composite literal to create a SLICE which holds 10 values of type int

	s := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	// range over the slice and print the values out

	for i, v := range s {
		fmt.Println(i, v)
	}
	// using format printing - print out type of the array
	fmt.Printf("%T\n", s)
}

//✨ Arrays are fixed-length sequences of items of the same type <--> Slices, on the other hand, are much more flexible, powerful, and convenient than arrays. Unlike arrays, slices can be resized using the built-in append function

//✨ Overall, slices are cleaner, more flexible, and less bug-prone than arrays, so you should prefer using them over arrays whenever possible.
