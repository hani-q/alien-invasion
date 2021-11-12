package main

import (
	"alien-invasion/city"
	"fmt"
)

func main() {

	c1 := city.City{Name: "Islamabad"}
	c2 := city.City{Name: "Lahore"}
	c3 := city.City{Name: "Faisalabad"}
	c4 := city.City{Name: "Karachi"}
	c5 := city.City{Name: "Quetta"}

	c1.North = city.Road{DirName: city.North, DestCity: &c2}
	c1.South = city.Road{DirName: city.South, DestCity: &c3}
	c1.East = city.Road{DirName: city.East, DestCity: &c4}
	c1.West = city.Road{DirName: city.West, DestCity: &c5}

	fmt.Println(c1)

}
