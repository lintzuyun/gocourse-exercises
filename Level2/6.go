package main

import "fmt"

func main() {

	// exercise #6
	// using iota, create 4 constants for the last 4 years and print the values.

	const (
		a = 2019 + iota
		b = 2019 + iota
		c = 2019 + iota
		d = 2019 + iota
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
