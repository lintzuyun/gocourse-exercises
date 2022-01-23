package main

import "fmt"

func main() {

	// exercise #8
	// create a program that uses a switch statement with no switch expression specified
	// ✨ if no switch expression specified -> default true ✨

	switch {
	case false:
		fmt.Println("The first one should not print")
	case true:
		fmt.Println("The second one should print")
	}

}
