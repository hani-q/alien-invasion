package structs

import (
	"testing"
)

func TestReverseDirectonNorth(t *testing.T) {

	var d Direction = North
	var wanted Direction = South
	changedD := d.ReverseDirecton()

	if d == South {
		t.Fatalf(`%v.ReverseDirecton() = %q, %v, want %v, error`, d, changedD,
			"Direction Not Reversed", wanted)
	}
}

func TestReverseDirectonSouth(t *testing.T) {

	var d Direction = South
	var wanted Direction = North
	changedD := d.ReverseDirecton()

	if d == North {
		t.Fatalf(`%v.ReverseDirecton() = %q, %v, want %v, error`, d, changedD,
			"Direction Not Reversed", wanted)
	}
}

func TestReverseDirectonEast(t *testing.T) {

	var d Direction = East
	var wanted Direction = West
	changedD := d.ReverseDirecton()

	if d == wanted {
		t.Fatalf(`%v.ReverseDirecton() = %q, %v, want %v, error`, d, changedD,
			"Direction Not Reversed", wanted)
	}
}

func TestReverseDirectonWest(t *testing.T) {

	var d Direction = West
	var wanted Direction = East
	changedD := d.ReverseDirecton()

	if d == wanted {
		t.Fatalf(`%v.ReverseDirecton() = %q, %v, want %v, error`, d, changedD,
			"Direction Not Reversed", wanted)
	}
}

func TestReverseStringDirectonNorth(t *testing.T) {

	var d = "noRth"
	wanted := "south"
	changedD := ReverseStringDirecton(d)

	if d == wanted {
		t.Fatalf(`ReverseStringDirecton(%v) = %q, %v, want %v, error`, d, changedD,
			"Direction Not Reversed", wanted)
	}
}

func TestReverseStringDirectonSouth(t *testing.T) {

	var d = "SoutH"
	wanted := "north"
	changedD := ReverseStringDirecton(d)

	if d == wanted {
		t.Fatalf(`ReverseStringDirecton(%v) = %q, %v, want %v, error`, d, changedD,
			"Direction Not Reversed", wanted)
	}
}

func TestReverseStringDirectonEast(t *testing.T) {

	var d = "EAST"
	wanted := "west"
	changedD := ReverseStringDirecton(d)

	if d == wanted {
		t.Fatalf(`ReverseStringDirecton(%v) = %q, %v, want %v, error`, d, changedD,
			"Direction Not Reversed", wanted)
	}
}

func TestReverseStringDirectonWest(t *testing.T) {

	var d = "weST"
	wanted := "east"
	changedD := ReverseStringDirecton(d)

	if d == wanted {
		t.Fatalf(`ReverseStringDirecton(%v) = %q, %v, want %v, error`, d, changedD,
			"Direction Not Reversed", wanted)
	}
}

// func (c *City) AddOccupant(a *Alien) bool {
// 	if len(c.occupants) < 2 {
// 		//Add if Alien not in the map
// 		if _, ok := c.occupants[a.Name]; !ok {
// 			c.occupants[a.Name] = a
// 			return true
// 		}

// 	}
// 	return false
// }
func TestAddAlienOccupantsToCity(t *testing.T) {

	var city = City{Name: "Testopolis", occupants: make(Occupants)}

	wanted := 2

	//Add more then 2 occupants and check status
	//Should be false
	city.AddOccupant(&Alien{Name: "ThandaKhan1"})
	city.AddOccupant(&Alien{Name: "ThandaKhan2"})
	city.AddOccupant(&Alien{Name: "ThandaKhan3"})

	count := len(city.occupants)

	if count != wanted {
		t.Fatalf(`city.AddOccupant(%v) = %v, %v, want %v, error`, wanted, count,
			"More then 2 Occupants added", wanted)
	}
}

func TestAddZeroAlienOccupantsToCity(t *testing.T) {

	var city = City{Name: "Testopolis", occupants: make(Occupants)}

	wanted := 0

	count := len(city.occupants)

	if count != wanted {
		t.Fatalf(`city.AddOccupant(%v) = %v, %v, want %v, error`, wanted, count,
			"More then 2 Occupants added", wanted)
	}
}

func TestAddRemoveAlienOccupantsToCity(t *testing.T) {

	var city = City{Name: "Testopolis", occupants: make(Occupants)}

	wanted := 1

	//Add more then 2 occupants and check status
	//Should be false
	city.AddOccupant(&Alien{Name: "ThandaKhan1"})
	city.AddOccupant(&Alien{Name: "ThandaKhan2"})

	city.RemoveOccupant("ThandaKhan2")

	count := len(city.occupants)

	if count != wanted {
		t.Fatalf(`city.RemoveOccupant(%v) = %v, %v, want %v, error`, wanted, count,
			"More then 2 Occupants added", wanted)
	}
}

func TestRemoveAlienOccupantsToCityWhenZeroAdded(t *testing.T) {

	var city = City{Name: "Testopolis", occupants: make(Occupants)}

	wanted := 0

	//Add more then 2 occupants and check status
	//Should be false

	city.RemoveOccupant("ThandaKhan1")

	count := len(city.occupants)

	if count != wanted {
		t.Fatalf(`city.RemoveOccupant(%v) = %v, %v, want %v, error`, wanted, count,
			"More then 2 Occupants added", wanted)
	}
}

func TestCountAlienOccupantsToCity(t *testing.T) {

	var city = City{Name: "Testopolis", occupants: make(Occupants)}

	wanted := 2

	//Add more then 2 occupants and check status
	//Should be false
	city.AddOccupant(&Alien{Name: "ThandaKhan1"})
	city.AddOccupant(&Alien{Name: "ThandaKhan2"})
	city.AddOccupant(&Alien{Name: "ThandaKhan3"})

	count := city.CountOccupants()

	if count != wanted {
		t.Fatalf(`city.AddOccupant(%v) = %v, %v, want %v, error`, wanted, count,
			"More then 2 Occupants added", wanted)
	}
}

func TestRandomNeighboutNoRoadCity(t *testing.T) {

	//Create Linked cities
	//Adding 0 Neigbour for each city1 and test
	var city1 = &City{Name: "Testopolis", occupants: make(Occupants)}
	var city2 = &City{Name: "Testabad", occupants: make(Occupants)}
	var _ = &City{Name: "Testmopoliton", occupants: make(Occupants), North: &Road{DirName: North, DestCity: city1},
		South: &Road{DirName: South, DestCity: city2}}

	neighbour, _ := city1.RandomNeighbour()

	if neighbour != nil {
		t.Fatalf(`city.RandomNeighbour(%v) = %v, %v, want %v, error`, nil, neighbour,
			"More then 2 Occupants added", nil)
	}
}
