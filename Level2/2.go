package main

import "fmt"

func main() {

	// exercise #2
	// using the following operators, write expressions and assign their values to variables:
	// == / <= / >= / != / < / >

	a := 10 == 10 // --> a:= (10 == 10) 可以加括號也可以不加
	b := 20 <= 10
	c := 20 >= 10
	d := 10 != 10
	e := 20 < 10
	f := 20 > 10

	fmt.Println(a, b, c, d, e, f)
}
