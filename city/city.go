package city

import (
	"fmt"
)

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case South:
		return "South"
	case East:
		return "East"
	case West:
		return "West"

	default:
		return "--"
	}
}

type City struct {
	Name                     string
	North, South, East, West Road
}

type Road struct {
	DirName  Direction
	DestCity *City
}

func (r Road) String() string {
	return fmt.Sprintf("%v =>%v", r.DirName, r.DestCity.Name)
}

func (c City) String() string {
	return fmt.Sprintf("======%v======\n\t%v\n%v\t%v\n\t%v\n=============", c.Name, c.North, c.East, c.West, c.South)
}
