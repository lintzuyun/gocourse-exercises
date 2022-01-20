package main

import "fmt"

func main() {

	// exercise #5
	// create a variable of type string using a raw string literal and print it.
	// ^ Raw string literals, delimited by backticks (back quotes), are interpreted literally. They can contain line breaks, and backslashes have no special meaning.

	string := `heeey
	I'm Joy
	"heeeey"
	This is raw string literal!`

	fmt.Println(string)
}
