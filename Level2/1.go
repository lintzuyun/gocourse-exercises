package main

import "fmt"

func main() {

	// exercise #1
	// write a program that prints a number in decimal(10進位), binary and hex(16進位)

	x := 32
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#x\n", x)

	// ^ can be written in one
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
}
