package structs

import (
	"fmt"
	"testing"
)

//Test loading of world map
//After loading we should have the same count as in map.txt file
func TestLoadWorldMap(t *testing.T) {

	//We know the test file has 10 cities
	wanted := 10
	testMapFilePath := "../test/world_test_ten.txt"
	LoadWorldMap(testMapFilePath)

	//XWorld the global world var should be updated
	cityCount := len(XWorld)

	if cityCount != wanted {
		t.Fatalf(`LoadWorldMap(%v) = %v, %v, want %v, error`, testMapFilePath, cityCount,
			"Cities in World struct should be 10", wanted)
	}
}

//Check when cities loaded have the proper directional links
//as stated in the map.txt file
func TestWorldCityLinks(t *testing.T) {

	//We know the test file has 10 cities
	// city 1 and 10 are linked
	// 1 is to east of 10 & 10 is to the west of 1
	testMapFilePath := "../test/world_test_ten.txt"

	LoadWorldMap(testMapFilePath)

	fmt.Println(XWorld)
	direction1 := XWorld["1"].West.DirName
	direction10 := XWorld["10"].East.DirName

	name1 := XWorld["1"].West.DestCity.Name
	name10 := XWorld["10"].East.DestCity.Name

	if direction1 != West {
		t.Fatalf(`LoadWorldMap(%v) %v , %v, want %v, error`, testMapFilePath, direction1, West,
			"city 10 should be West")
	}

	if direction10 != East {
		t.Fatalf(`LoadWorldMap(%v) %v, %v, want %v, error`, testMapFilePath, direction10, East,
			"city 1 should be East")
	}

	if name1 != "10" {
		t.Fatalf(`LoadWorldMap(%v) %v, %v, want %v, error`, testMapFilePath, name1, "10",
			"city 1 should have neighbour named 10")
	}

	if name10 != "1" {
		t.Fatalf(`LoadWorldMap(%v) %v, %v, want %v, error`, testMapFilePath, name10, "1",
			"city 10 should have neighbout name 1")
	}
}

//Check if when a city is deleteed its neighbours should aslo forget about it
func TestDeleteCityAndLinks(t *testing.T) {

	//We know the test file has 10 cities
	// city 1 and 10 are linked
	// 1 is to east of 10 & 10 is to the west of 1
	testMapFilePath := "../test/world_test_ten.txt"

	LoadWorldMap(testMapFilePath)

	fmt.Println(XWorld)
	city1 := XWorld["1"]
	city10 := XWorld["10"]

	city10Name := city10.Name

	//Delete City 10 and see if City 1 has 1 neighbout left or not
	XWorld.DeleteCity(city10Name)

	if city1.East != nil {
		t.Fatalf(`DeleteCity(%v) %v, %v, want %v, error`, "city10", nil, city10Name,
			"city 1 East neigh 10 was deleted and should be nil")
	}

}

//Test when aliens are placed in a world each one is assigned to a single city
func TestPlaceAliensCount(t *testing.T) {

	//We know the test file has 10 cities
	testMapFilePath := "../test/world_test_ten.txt"

	LoadWorldMap(testMapFilePath)

	alienCount := 10
	XWorld.PlaceTheAliens(alienCount)
	placedAliens := 0

	for _, cityData := range XWorld {
		placedAliens += cityData.CountOccupants()
		if cityData.CountOccupants() > 1 {
			t.Fatalf(`PlaceTheAliens(%v) = %v, %v, want %v, error`, alienCount, placedAliens,
				"Aliens placed in world should be 10", alienCount)
		}
	}

	if alienCount != placedAliens {
		t.Fatalf(`PlaceTheAliens(%v) = %v, %v, want %v, error`, alienCount, placedAliens,
			"Aliens placed in world should be 10", alienCount)
	}
}

//Test when aliens are placed in a world each one is not duplicated
func TestPlaceAliensDuplication(t *testing.T) {

	//We know the test file has 10 cities
	testMapFilePath := "../test/world_test_ten.txt"

	LoadWorldMap(testMapFilePath)

	alienCount := 10
	XWorld.PlaceTheAliens(alienCount)
	placedAlienNames := make([]string, alienCount)

	for _, cityData := range XWorld {
		for k := range cityData.occupants {
			if checkAlienInArray(placedAlienNames, k) {
				t.Fatalf(`PlaceTheAliens(%v) = %v, %v, want %v, error`, alienCount, true, false,
					"Aliens placed in world should be unique")
			}
		}
	}

}

func checkAlienInArray(s []string, name string) bool {
	for _, n := range s {
		if n == name {
			return true
		}
	}

	return false
}
