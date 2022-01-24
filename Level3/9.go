package main

import "fmt"

func main() {

	// exercise #9
	// create a program that uses a switch statement with the switch expression specified as a variabel of TYPE string with the IDENTIFIER "favSport"

	favSport := "badminton"
	switch favSport {
	case "baseball":
		fmt.Println("I like baseball")
	case "badminton":
		fmt.Println("I like badminton")
	default:
		fmt.Println("I don't like sports")
	}

}
