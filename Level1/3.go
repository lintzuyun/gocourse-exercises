package main

import "fmt"

// exercise #3 -1 (package scope!)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	// exercise #3 -2

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z) // \t = the tab character
	fmt.Println(s)
}
