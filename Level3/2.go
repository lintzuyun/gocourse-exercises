package main

import "fmt"

// Print every rune code point of the uppercase alphabet three times
// https://en.wikipedia.org/wiki/ASCII
// ASCII A = 65 / Z = 90

func main() {
	// exercise #2
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
