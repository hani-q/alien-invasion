package structs

import (
	"fmt"
	"sync"
)

type Direction string

const (
	X     Direction = "X"
	North Direction = "North"
	South Direction = "South"
	East  Direction = " East"
	West  Direction = " West"
)

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
	Invader                  *Alien
	mu                       sync.Mutex
}

func (c *City) RandomNeighbour() (*City, string) {

	availableNeighbours := make(map[string]*City)

	//Make a slice of Neighs that are not NIL and dont match Current City Name
	//or the Previously Visited City Name
	if c.North != nil {
		if c.North.DestCity != nil {
			if c.Name != c.North.DestCity.Name {
				availableNeighbours["North"] = c.North.DestCity
			}
		}
	}
	if c.South != nil {
		if c.South.DestCity != nil {
			if c.Name != c.South.DestCity.Name {
				availableNeighbours["South"] = c.South.DestCity
			}
		}

	}

	if c.East != nil {
		if c.East.DestCity != nil {
			if c.Name != c.East.DestCity.Name {
				availableNeighbours["East"] = c.East.DestCity
			}
		}
	}

	if c.West != nil {
		if c.West.DestCity != nil {
			if c.Name != c.West.DestCity.Name {
				availableNeighbours["West"] = c.West.DestCity
			}
		}
	}

	if len(availableNeighbours) > 0 {
		for k, v := range availableNeighbours {
			return v, k
		}
	}

	return nil, ""

}

func (c *City) String() string {
	return fmt.Sprintf("%v %v %v %v %v", c.Name, c.North.getRoadName(), c.East.getRoadName(), c.West.getRoadName(), c.South.getRoadName())
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
