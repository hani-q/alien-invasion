package structs

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
		return "north"
	case South:
		return "south"
	case East:
		return "east"
	case West:
		return "west"
	default:
		return ""
	}
}

func (d Direction) ReverseDirecton() Direction {
	switch d {
	case X:
		return X
	case North:
		return South
	case South:
		return East
	case East:
		return West
	case West:
		return East
	default:
		return X
	}
}

func ReverseStringDirecton(dir string) string {
	switch dir {
	case "north":
		return "south"
	case "east":
		return "west"
	case "west":
		return "east"
	case "south":
		return "north"
	default:
		return ""
	}
}

type City struct {
	Name                     string
	North, South, East, West *Road
	Invaders                 [2]*Alien
}

func (c City) String() string {
	return fmt.Sprintf("%v %v %v %v %v [%v, %v]", c.Name, c.North.getRoadName(), c.East.getRoadName(), c.West.getRoadName(), c.South.getRoadName(), c.Invaders[0], c.Invaders[1])
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

func (r *Road) getRoadName() string {
	if r == nil {
		return ""
	}

	return fmt.Sprintf("%v=%v", r.DirName, r.DestCity.getCityName())
}
