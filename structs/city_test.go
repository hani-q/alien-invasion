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

//Test when only road leads out of city
func TestRandomNeighbourOneRoad(t *testing.T) {

	//Create Linked cities
	var city1 = &City{Name: "Testopolis"}
	var city2 = &City{Name: "Testabad"}
	var city3 = &City{Name: "Testmopoliton", North: &Road{DirName: North, DestCity: city1},
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
	var city1 = &City{Name: "Testopolis"}
	var city2 = &City{Name: "Testabad"}
	var city3 = &City{Name: "Testmopoliton", North: &Road{DirName: North, DestCity: city1},
		South: &Road{DirName: South, DestCity: city2}}

	neighbour, _ := city3.RandomNeighbour()

	if neighbour == nil {
		t.Fatalf(`city.RandomNeighbour(%v) = %v, %v, want %v, error`, 1, neighbour,
			"1 Neighbour should have been found", 1)
	}
}
