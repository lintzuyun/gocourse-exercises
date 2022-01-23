package main

import "fmt"

func main() {

	// exercise #4
	// Print out the years you have been alive with syntax -> for {}

	bd := 1991
	for {
		if bd > 2022 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
