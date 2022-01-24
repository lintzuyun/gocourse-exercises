package main

import "fmt"

func main() {

	// exercise #3
	// use SLICING to create new slices
	s := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	fmt.Println(s[:5])
	fmt.Println(s[5:])
	fmt.Println(s[2:7])
	fmt.Println(s[1:6])
	fmt.Println()
}
