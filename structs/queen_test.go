package structs

import (
	"testing"
)

// Test when aliens are with less count
func TestLayEggsAliensZeroPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	alienCount := 0
	queen := Queen{Children: make(map[string]*Alien), QueenChan: make(chan AlienLanguage)}

	//We know the test file has 10 cities
	testMapFilePath := "../test/world_test_ten.txt"

	xWorld := LoadWorldMap(testMapFilePath)

	//Spew the Queen Mothers Eggs the Entire X-World
	queen.LayEggs(alienCount, xWorld)

}
