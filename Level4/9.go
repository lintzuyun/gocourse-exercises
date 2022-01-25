package main

import "fmt"

func main() {

	// exercise #9

	/*
	  Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop

	*/

	m := map[string][]string{
		`bond_james`: []string{`Shaken, not stirred`, `Martinis`, `Women`}, `moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`}, `no_dr`: []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	//✨
	m[`joy_lin`] = []string{`Netflix`, `Kimchi`, `Popcorn`}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
