package structs

import (
	"testing"
)

//Test if a direction is properly reversed
func TestReverseDirectonNorth(t *testing.T) {

	var d Direction = North
	var wanted Direction = South
	changedD := d.ReverseDirecton()

	if d == South {
		t.Fatalf(`%v.ReverseDirecton() = %q, %v, want %v, error`, d, changedD,
			"Direction Not Reversed", wanted)
	}
}

//Test if a direction is properly reversed
func TestReverseDirectonSouth(t *testing.T) {

	var d Direction = South
	var wanted Direction = North
	changedD := d.ReverseDirecton()

	if d == North {
		t.Fatalf(`%v.ReverseDirecton() = %q, %v, want %v, error`, d, changedD,
			"Direction Not Reversed", wanted)
	}
}

//Test if a direction is properly reversed
func TestReverseDirectonEast(t *testing.T) {

	var d Direction = East
	var wanted Direction = West
	changedD := d.ReverseDirecton()

	if d == wanted {
		t.Fatalf(`%v.ReverseDirecton() = %q, %v, want %v, error`, d, changedD,
			"Direction Not Reversed", wanted)
	}
}

//Test if a direction is properly reversed
func TestReverseDirectonWest(t *testing.T) {

	var d Direction = West
	var wanted Direction = East
	changedD := d.ReverseDirecton()

	if d == wanted {
		t.Fatalf(`%v.ReverseDirecton() = %q, %v, want %v, error`, d, changedD,
			"Direction Not Reversed", wanted)
	}
}

//Test if a direction is properly reversed
func TestReverseStringDirectonNorth(t *testing.T) {

	var d = "noRth"
	wanted := "south"
	changedD := ReverseStringDirecton(d)

	if d == wanted {
		t.Fatalf(`ReverseStringDirecton(%v) = %q, %v, want %v, error`, d, changedD,
			"Direction Not Reversed", wanted)
	}
}

//Test if a direction is properly reversed
func TestReverseStringDirectonSouth(t *testing.T) {

	var d = "SoutH"
	wanted := "north"
	changedD := ReverseStringDirecton(d)

	if d == wanted {
		t.Fatalf(`ReverseStringDirecton(%v) = %q, %v, want %v, error`, d, changedD,
			"Direction Not Reversed", wanted)
	}
}

//Test if a direction is properly reversed
func TestReverseStringDirectonEast(t *testing.T) {

	var d = "EAST"
	wanted := "west"
	changedD := ReverseStringDirecton(d)

	if d == wanted {
		t.Fatalf(`ReverseStringDirecton(%v) = %q, %v, want %v, error`, d, changedD,
			"Direction Not Reversed", wanted)
	}
}

//Test if a direction is properly reversed
func TestReverseStringDirectonWest(t *testing.T) {

	var d = "weST"
	wanted := "east"
	changedD := ReverseStringDirecton(d)

	if d == wanted {
		t.Fatalf(`ReverseStringDirecton(%v) = %q, %v, want %v, error`, d, changedD,
			"Direction Not Reversed", wanted)
	}
}

//Test if Adding occupants to city Map
//it doesnt exceed the limit of 2
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

//Test when 0 occpuants is passed as param
func TestAddZeroAlienOccupantsToCity(t *testing.T) {

	var city = City{Name: "Testopolis", occupants: make(Occupants)}

	wanted := 0

	count := len(city.occupants)

	if count != wanted {
		t.Fatalf(`city.AddOccupant(%v) = %v, %v, want %v, error`, wanted, count,
			"More then 2 Occupants added", wanted)
	}
}

//Test addition and then removal
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

//Test removal wihout addition
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

//Test count of all aliens added to city should not exceed 2
//even when adding more then 2
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

//Test if neghbout has no roads going out
func TestRandomNeighbourNoRoad(t *testing.T) {
	//Create Linked cities
	var city1 = &City{Name: "Testopolis", occupants: make(Occupants)}
	var city2 = &City{Name: "Testabad", occupants: make(Occupants)}
	var _ = &City{Name: "Testmopoliton", occupants: make(Occupants), North: &Road{DirName: North, DestCity: city1},
		South: &Road{DirName: South, DestCity: city2}}

	neighbour, _ := city1.RandomNeighbour()

	if neighbour != nil {
		t.Fatalf(`city.RandomNeighbour(%v) = %v, %v, want %v, error`, nil, neighbour,
			"Neighbour should not have been found", nil)
	}
}

//Test when only road leads out of city
func TestRandomNeighbourOneRoad(t *testing.T) {

	//Create Linked cities
	var city1 = &City{Name: "Testopolis", occupants: make(Occupants)}
	var city2 = &City{Name: "Testabad", occupants: make(Occupants)}
	var city3 = &City{Name: "Testmopoliton", occupants: make(Occupants), North: &Road{DirName: North, DestCity: city1},
		South: &Road{DirName: South, DestCity: city2}}

	neighbour, _ := city3.RandomNeighbour()

	if neighbour == nil {
		t.Fatalf(`city.RandomNeighbour(%v) = %v, %v, want %v, error`, 1, neighbour,
			"1 Neighbour should have been found", 1)
	}
}

//Test when 2 roads lead out of city
func TestRandomNeighbourTwoRoad(t *testing.T) {

	//Create Linked cities
	var city1 = &City{Name: "Testopolis", occupants: make(Occupants)}
	var city2 = &City{Name: "Testabad", occupants: make(Occupants)}
	var city3 = &City{Name: "Testmopoliton", occupants: make(Occupants), North: &Road{DirName: North, DestCity: city1},
		South: &Road{DirName: South, DestCity: city2}}

	neighbour, _ := city3.RandomNeighbour()

	if neighbour == nil {
		t.Fatalf(`city.RandomNeighbour(%v) = %v, %v, want %v, error`, 1, neighbour,
			"1 Neighbour should have been found", 1)
	}
}

//Test when roads leads to a city with 2 occupants
func TestRandomNeighbourOneRoadTwoOccupants(t *testing.T) {

	//Create Linked cities
	var city1 = &City{Name: "Testopolis", occupants: make(Occupants)}
	city1.AddOccupant(&Alien{Name: "1"})
	city1.AddOccupant(&Alien{Name: "2"})

	var city2 = &City{Name: "Testmopoliton", occupants: make(Occupants), North: &Road{DirName: North, DestCity: city1}}

	neighbour, _ := city2.RandomNeighbour()

	if neighbour != nil {
		t.Fatalf(`city.RandomNeighbour(%v) = %v, %v, want %v, error`, nil, neighbour,
			"0 Neighbour should have been found", nil)
	}
}
