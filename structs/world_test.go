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
	xWorld := LoadWorldMap(testMapFilePath)

	cityCount := xWorld.GetCityCount()

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

	xWorld := LoadWorldMap(testMapFilePath)

	fmt.Println(xWorld)
	direction1 := xWorld.Data["1"].West.DirName
	direction10 := xWorld.Data["10"].East.DirName

	name1 := xWorld.Data["1"].West.DestCity.Name
	name10 := xWorld.Data["10"].East.DestCity.Name

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

	xWorld := LoadWorldMap(testMapFilePath)

	fmt.Println(xWorld)
	city1 := xWorld.Data["1"]
	city10 := xWorld.Data["10"]

	city10Name := city10.Name

	//Delete City 10 and see if City 1 has 1 neighbout left or not
	xWorld.DeleteCity(city10Name)

	if city1.East != nil {
		t.Fatalf(`DeleteCity(%v) %v, %v, want %v, error`, "city10", nil, city10Name,
			"city 1 East neigh 10 was deleted and should be nil")
	}

}

//Test when aliens are placed in a world each one is assigned to a single city
func TestLayEggsCount(t *testing.T) {

	//We know the test file has 10 cities
	testMapFilePath := "../test/world_test_ten.txt"

	alienCount := 10
	xWorld := LoadWorldMap(testMapFilePath)
	queen := Queen{Children: make(map[string]*Alien), QueenChan: make(chan AlienLanguage)}

	//Spew the Queen Mothers Eggs the Entire X-World
	queen.LayEggs(alienCount, xWorld)

	placedAliens := 0

	for _, cityData := range xWorld.Data {
		if cityData.Occupant != nil {
			placedAliens++
		}
	}

	if alienCount != placedAliens {
		t.Fatalf(`PlaceTheAliens(%v) = %v, %v, want %v, error`, alienCount, placedAliens,
			"Aliens placed in world should be 10", alienCount)
	}
}

// //Test when aliens are placed in a world each one is not duplicated
func TestLayEggsDuplication(t *testing.T) {

	//We know the test file has 10 cities
	testMapFilePath := "../test/world_test_ten.txt"

	xWorld := LoadWorldMap(testMapFilePath)

	alienCount := 10
	queen := Queen{Children: make(map[string]*Alien), QueenChan: make(chan AlienLanguage)}

	//Spew the Queen Mothers Eggs the Entire X-World
	queen.LayEggs(alienCount, xWorld)

	placedAlienNames := make([]string, alienCount)

	for _, cityData := range xWorld.Data {
		if cityData.Occupant != nil {
			alieName := cityData.Occupant.Name
			if checkAlienInArray(placedAlienNames, alieName) {
				t.Fatalf(`PlaceTheAliens(%v) = %v, %v, want %v, error`, alienCount, true, false,
					"Aliens placed in world should be unique")
			} else {
				placedAlienNames = append(placedAlienNames, alieName)
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
