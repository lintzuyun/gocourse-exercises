package main

import "fmt"

// exercise #3
type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	truck1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		fourWheel: true,
	}

	sedan1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: false,
	}

	fmt.Println(truck1)
	fmt.Println(truck1.doors)
	fmt.Println(truck1.fourWheel)
	fmt.Println(sedan1)
	fmt.Println(sedan1.doors)
	fmt.Println(sedan1.luxury)
}
