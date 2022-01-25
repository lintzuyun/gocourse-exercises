package main

import "fmt"

// exercise #2
// Take the code from the previous exercise, then store the values of type person in a MAP with the key of last name. Access each value in the map. Print out the values, ranging over the slice.

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

	// fmt.Println(p1.firstName)
	// fmt.Println(p1.lastName)
	// for i, v := range p1.favoriteIceCreamFlavors {
	// 	fmt.Println("\t", i, v)
	// }

	// fmt.Println(p2.firstName)
	// fmt.Println(p2.lastName)
	// for i, v := range p2.favoriteIceCreamFlavors {
	// 	fmt.Println("\t", i, v)
	// }

	//✨把資料存進map裡
	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	fmt.Println(m)

	//不需要key的時候可以用_取代

	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, val := range v.favoriteIceCreamFlavors {
			fmt.Println(i, val)
		}
	}
}
