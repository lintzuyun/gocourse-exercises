package main

import "fmt"

// exercise #5 -1
type joy int

var x joy
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x) // (1) Printf = format print (2) \n = line break
	x = 1234
	fmt.Println(x)

	// exercise #5 -2
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
