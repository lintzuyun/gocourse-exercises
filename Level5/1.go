package main

import "fmt"

// exercise #1

type person struct {
	firstName               string
	lastName                string
	favoriteIceCreamFlavors []string
}

func main() {

	p1 := person{
		firstName:               "joy",
		lastName:                "lin",
		favoriteIceCreamFlavors: []string{"vanilla", "strawberry"},
	}

	p2 := person{
		firstName:               "craig",
		lastName:                "bickford",
		favoriteIceCreamFlavors: []string{"chocolate", "cream", "mocha"},
	}

	// fmt.Println(p1, p2)

	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	for i, v := range p1.favoriteIceCreamFlavors {
		fmt.Println("\t", i, v)
	}

	fmt.Println(p2.firstName)
	fmt.Println(p2.lastName)
	for i, v := range p2.favoriteIceCreamFlavors {
		fmt.Println("\t", i, v)
	}
}
