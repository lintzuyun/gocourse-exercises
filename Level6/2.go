package main

import "fmt"

func main() {

	// exercise #2

	// for foo()
	x := []int{1, 2, 3, 4, 5, 6, 7} //pass in a value of type []int into your func (unfurl the []int)
	y := foo(x...)
	fmt.Println(y)

	// for bar()
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := bar(a)
	fmt.Println(b)

}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(y []int) int {
	total := 0
	for _, v := range y {
		total += v
	}

	return total
}
