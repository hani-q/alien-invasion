package city

import (
	"fmt"
)

type Direction int

const (
	X Direction = iota
	North
	South
	East
	West
)

func (d Direction) String() string {
	switch d {
	case X:
		return ""
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

func (c *City) getCityName() string {
	if c == nil {
		return "XXXX"
	}

	return c.Name
}

func (r Road) String() string {
	if r.DirName == X {
		return ""
	}

	return fmt.Sprintf("%v =>%v", r.DirName, r.DestCity.getCityName())
}

func (c City) String() string {

	return fmt.Sprintf("======%v======\n%v\n%v%v\n%v\n=====================", c.Name, c.North, c.East, c.West, c.South)
}
