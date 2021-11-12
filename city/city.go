package city

import "fmt"

type City struct {
	Name       string   //Name of city
	Directions []string // Bridges leading out from city
}

func (c *City) String() string {
	return fmt.Sprintf("%v\n", (&c).Name)
}
