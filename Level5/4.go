package main

import "fmt"

func main() {

	// exercise #4
	// Create and use an anonymous struct.

	s := struct {
		first     string
		friends   map[string]int // map
		favDrinks []string       // slice
	}{
		first: "Joy",
		friends: map[string]int{
			"B": 22,
			"C": 33,
			"D": 44,
		},
		favDrinks: []string{
			"Coke",
			"Juice",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	// map - key & value
	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	// slice - index & value
	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}
}
