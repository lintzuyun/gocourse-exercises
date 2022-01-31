package main

import "fmt"

func main() {

	// exercise #3
	defer first()
	second()
}

func first() {
	fmt.Println("This should come second")

}

func second() {
	fmt.Println("This should come first")
}
