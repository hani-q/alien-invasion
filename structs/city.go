package structs

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
)

type Direction string

//Cardinal Directions Enum
const (
	X     Direction = "X"
	North Direction = "North"
	South Direction = "South"
	East  Direction = " East"
	West  Direction = " West"
)

//Will return the Opposite/Reverse Direction of a given
//Direction... N<->S ... E<->W
func (d Direction) ReverseDirecton() Direction {
	switch d {
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

//Helper function to reverse the given direction in string format
func ReverseStringDirecton(dir string) string {
	dir = strings.ToLower(dir)
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
	Occupant                 *Alien //internal so that nobody canadd more then 2 Aliens
	mu                       sync.Mutex
}

//Returns a random neighbour of the City.
//only returns a non nill value if no Roads Lead out of city
//or if the available neighbours are not full
//Nobody should move to a city that has 2 Aliens fighting it out
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
		for k, neighbour := range availableNeighbours {
			return neighbour, k
		}
	}

	return nil, ""

}

//Prints City in input map file format
//Bar south=Foo west=Bee
func (c *City) String() string {
	var roadData string
	if c.North != nil {
		roadData += c.North.getRoadName()
	}
	if c.South != nil {
		roadData += c.South.getRoadName()
	}
	if c.East != nil {
		roadData += c.East.getRoadName()
	}
	if c.West != nil {
		roadData += c.West.getRoadName()
	}

	//Remove any double spaces
	space := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	roadData = space.ReplaceAllString(roadData, " ")

	return fmt.Sprintf("%v%s\n", c.Name, roadData)
}

//Alternate Print used in Debugging.
//Shows a city, its roads and its aliens
func (c *City) CityPrint() string {
	return fmt.Sprintf("Occupant: %v, Roads: %v %v %v %v\n", c.Occupant, c.North.getRoadName(), c.East.getRoadName(), c.West.getRoadName(), c.South.getRoadName())
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

	return fmt.Sprintf(" %v=%v", r.DirName, r.DestCity.getCityName())
}

// //Instanitiate the aliens with count provided in cli Args
// //Each alien will be placed in a loving and caring city
// //It will be made sure that no other alien will be present in the
// //same city
// func SpawnAliens(alienCount int) []*Alien {

// 	if alienCount < 2 {
// 		msg := "Error: Alient count cannot be less then 2"
// 		log.Error(msg)
// 		panic(msg)
// 	}

// 	//Get fancy alien names from this NameGenerator library
// 	seed := time.Now().UTC().UnixNano()
// 	nameGenerator := namegenerator.NewNameGenerator(seed)

// 	//Keep a track of Spawned Aliens
// 	//This struct will be used by start simulation function to
// 	//Launch the aliens
// 	var result []*Alien = make([]*Alien, alienCount)
// 	for i := 0; i < alienCount; i++ {
// 		result[i] = &Alien{Name: util.Capitalise(nameGenerator.Generate())}
// 	}
// 	return result
// }
