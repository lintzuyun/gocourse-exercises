package main

import "fmt"

func main() {

	// exercise #1
	// using a composite literal to create an array which holds 5 values of type int

	a := [5]int{10, 11, 12, 13, 14}

	// range over the array and print the values out

	for i, v := range a {
		fmt.Println(i, v)
	}
	// using format printing - print out type of the array
	fmt.Printf("%T\n", a)
}
