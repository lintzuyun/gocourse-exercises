package main

import "fmt"

func main() {

	// exercise #8
	/*
			● Create a func which returns a func
		    ● assign the returned func to a variable
		    ● call the returned func
	*/

	f := function()
	fmt.Println(f)
	fmt.Println(f())
}

func function() func() int {
	return func() int {
		return 42
	}
}
